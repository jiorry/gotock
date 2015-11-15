package tdx

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"

	"github.com/kere/gos/db"
)

var threadNum int = 50

type IStore interface {
	ReadData(string) (db.DataSet, error)
	fill(string, db.DataSet) error
	DealRead(string) db.DataSet
	Parent() IStore
}

type StoreBase struct {
	table, prefix, ext, folder string
	parent                     IStore
}

func (a *StoreBase) Parent() IStore {
	return a.parent
}

func (a *StoreBase) DealRead(code string) db.DataSet {
	var result db.DataSet
	var err error

	result, err = a.Parent().ReadData(path.Join(a.folder, fmt.Sprint(a.prefix, code, a.ext)))
	fmt.Println("DealRead", code, len(result))

	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	return result
}

func dealitem(prefix, folder, code, ext string, controlChan chan string) {
	var a IStore
	switch ext {
	case ".day":
		a = NewStoreDaily(prefix, folder)
	case ".lc5":
		a = NewStore5m(prefix, folder)
	case ".lc1":
		a = NewStore1m(prefix, folder)
	}

	if err := a.Parent().fill(code, a.DealRead(code)); err != nil {
		fmt.Println("error:", err)
		controlChan <- ""
	} else {
		controlChan <- code
	}
	a = nil
}

func ScanAndStore(prefix, folder, ext string) error {
	codelist, err := ScanDirGetAllFileName(folder, ext)
	if err != nil {
		return err
	}
	fmt.Println("start scan and store", prefix, ext, len(codelist), "files")

	os.Mkdir(path.Join(folder, "bk"), os.ModeDir)

	var size = len(codelist)
	var finishRead = 0
	var worked = 0
	var controlChan chan string = make(chan string, threadNum)
	var runningChan chan int = make(chan int, threadNum)

	go func() {
		for i := 0; i < size; i++ {
			worked++
			runningChan <- 1
			go dealitem(prefix, folder, codelist[i], ext, controlChan)
		}
	}()

	for {
		<-runningChan
		code := <-controlChan
		if code != "" {
			name := fmt.Sprint(prefix, code, ext)
			err := os.Rename(path.Join(folder, name), path.Join(folder, "bk", name))
			if err != nil {
				fmt.Println("error: ", err)
			}

			fmt.Println("end:", code, ext, "===================")
		} else {
			fmt.Println("failed:", "********************")
		}

		finishRead++
		if finishRead == worked {
			break
		}

		if math.Mod(float64(finishRead), float64(threadNum)) == 0 {
			fmt.Println("runtime.GC:", "****************************************")
			runtime.GC()
		}
	}

	return nil
}

func (a *StoreBase) prepareStock(prefix, code string) int64 {
	q := db.NewQueryBuilder("stock").Where("code=? and ctype=?", code, prefix).Select("id")
	row, err := q.QueryOne()
	if err != nil {
		fmt.Println(err)
		return int64(-1)
	}
	if row != nil {
		return row.GetInt64("id")
	}

	insert := db.NewInsertBuilder("stock")
	data := db.DataRow{}
	data["code"] = code
	data["ctype"] = a.prefix
	_, err = insert.Insert(data)
	if err != nil {
		fmt.Println(err)
		return int64(-1)
	}

	row, err = q.QueryOne()
	if err != nil {
		fmt.Println(err)
		return int64(-1)
	}

	if row.Empty() {
		fmt.Println("empty row", prefix, code)
		return int64(-1)
	}

	return row.GetInt64("id")
}
