package tdx

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/kere/gos/db"
)

type StoreDaily struct {
	StoreBase
}

func NewStoreDaily(prefix, folder string) *StoreDaily {
	s := &StoreDaily{}
	s.StoreBase.table = "stock_daily"
	s.StoreBase.prefix = prefix
	s.StoreBase.ext = ".day"
	s.StoreBase.folder = folder
	s.parent = s
	return s
}

func (a *StoreDaily) fill(code string, result db.DataSet) error {
	if len(result) == 0 {
		return nil
	}

	id := a.prepareStock(a.StoreBase.prefix, code)
	if id < 0 {
		return fmt.Errorf("stock not found %s %s", a.StoreBase.prefix, code)
	}

	insert := db.NewInsertBuilder(a.table)
	// exists := db.NewExistsBuilder(a.table)
	fmt.Println("begin write records:", a.prefix, code, " total: ", len(result))

	date := ""
	rows := db.DataSet{}
	k := 0
	for _, item := range result {
		item["stock_id"] = id
		date = item.GetString("date")
		// if exists.Where("stock_id=? and date=?", id, date).Exists() {
		// 	fmt.Println("skip:", code, id, date)
		// 	continue
		// }

		rows = append(rows, item)
		// fmt.Println("insert:", code, id, date)

		if k == 20 {
			k = 0
			insert.InsertM(rows)
			fmt.Println(code, id, date, "---- InsertM ----")
			rows = db.DataSet{}
		}
		k++
	}
	if len(rows) > 0 {
		insert.InsertM(rows)
		fmt.Println("--------------- InsertM ---------------")
	}
	return nil
}

func (a *StoreDaily) ReadData(name string) (db.DataSet, error) {
	result := db.DataSet{}

	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return result, err
	}

	defer file.Close()
	var fields []string = []string{"date", "open", "high", "low", "close", "amount", "volumn"}

	var b []byte = make([]byte, 4)
	var data db.DataRow

	var x int32
	var xf float32

	var n = 8
	var i int = 0

	for true {
		if i == n {
			i = 0
		}
		_, err = file.Read(b)
		if err != nil {
			break
		}

		b_buf := bytes.NewBuffer(b)
		// binary.Read(b_buf, binary.BigEndian, &x)

		switch i {
		case 0:
			err = binary.Read(b_buf, binary.LittleEndian, &x)
			if err != nil {
				return nil, err
			}
			data = db.DataRow{}
			s := fmt.Sprintf("%d", x)
			data[fields[0]] = fmt.Sprintf("%s-%s-%s", s[0:4], s[4:6], s[6:8])

		case 1, 2, 3, 4:
			err = binary.Read(b_buf, binary.LittleEndian, &x)
			if err != nil {
				return nil, err
			}
			data[fields[i]] = float64(x) / 100

		case 5:
			err = binary.Read(b_buf, binary.LittleEndian, &xf)
			if err != nil {
				return nil, err
			}
			data[fields[i]] = fmt.Sprintf("%.0f", xf)

		case 6:
			err = binary.Read(b_buf, binary.LittleEndian, &x)
			if err != nil {
				return nil, err
			}
			data[fields[i]] = x

		case 7:
			result = append(result, data)

		}

		i++
	}

	return result, nil
}
