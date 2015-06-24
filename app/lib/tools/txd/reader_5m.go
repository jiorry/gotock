package txd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/kere/gos/db"
	"math"
	"os"
)

type Store5m struct {
	StoreBase
}

func NewStore5m(prefix, folder string) *Store5m {
	s := &Store5m{}
	s.StoreBase.table = "stock_5m"
	s.StoreBase.prefix = prefix
	s.StoreBase.ext = ".lc5"
	s.StoreBase.folder = folder
	s.parent = s
	return s
}

func (a *Store5m) fill(code string, result db.DataSet) error {
	if len(result) == 0 {
		return nil
	}

	id := a.prepareStock(a.StoreBase.prefix, code)
	if id < 0 {
		return fmt.Errorf("stock not found %s %s", a.StoreBase.prefix, code)
	}

	// tx := &db.Tx{}
	insert := db.NewInsertBuilder(a.StoreBase.table)
	// exists := db.NewExistsBuilder(a.StoreBase.table)
	fmt.Println("begin write records:", a.StoreBase.prefix, code, " total: ", len(result))

	rows := db.DataSet{}
	date := ""
	seq := 0
	k := 0

	for _, item := range result {
		item["stock_id"] = id
		if item.IsSet("other") {
			delete(item, "other")
		}

		if date != item.GetString("date") {
			date = item.GetString("date")
			seq = 0
		}

		seq++

		item["seq"] = seq

		// if exists.Where("stock_id=? and date=? and seq=?", id, date, seq).Exists() {
		// 	fmt.Println("skip", code, id, date, seq)
		// 	continue
		// }

		// insert.TxInsert(tx, item)
		rows = append(rows, item)
		// fmt.Println("insert", code, id, date, seq)

		if k == 20 {
			k = 0

			_, err := insert.InsertM(rows)
			if err != nil {
				return err
			}

			// fmt.Println(code, id, date, seq, "---- InsertM ----")
			rows = db.DataSet{}
		}
		k++
	}

	if len(rows) > 0 {
		_, err := insert.InsertM(rows)
		if err != nil {
			return err
		}
		fmt.Println("--------------- InsertM ---------------")
	}

	fmt.Println("end write")
	return nil
}

func (a *Store5m) ReadData(name string) (db.DataSet, error) {
	result := db.DataSet{}

	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return result, err
	}
	defer file.Close()

	var fields []string = []string{"date", "open", "high", "low", "close", "amount", "volumn", "other", "a8", "a9", "a10", "a11", "a12", "a12", "a14", "a15", "a16"}

	var b []byte
	var data db.DataRow
	var x_u16 uint16
	var x_u32 uint32
	var xf float32
	var n = 8
	var i int = 0

	for true {
		if i == n {
			i = 0
		}

		b = make([]byte, 4)
		_, err = file.Read(b)
		if err != nil {
			break
		}

		b_buf := bytes.NewBuffer(b)

		switch i {
		case 0:
			data = db.DataRow{}
			err = binary.Read(b_buf, binary.LittleEndian, &x_u16)
			if err != nil {
				return result, nil
			}
			year := float64(x_u16)/2048.00 + 2004
			mod := math.Mod(float64(x_u16), 2048)
			month := math.Floor(mod / 100)
			day := math.Mod(mod, 100.00)

			data[fields[i]] = fmt.Sprintf("%d-%02d-%02d", int(year), int(month), int(day))

		case 5:
			err = binary.Read(b_buf, binary.LittleEndian, &xf)
			if err != nil {
				return result, nil
			}
			data[fields[i]] = uint32(xf)

		case 6:
			err = binary.Read(b_buf, binary.LittleEndian, &x_u32)
			if err != nil {
				return result, nil
			}
			data[fields[i]] = x_u32 / 100

		default:
			err = binary.Read(b_buf, binary.LittleEndian, &xf)
			if err != nil {
				return result, nil
			}
			data[fields[i]] = xf

		case n - 1:
			err = binary.Read(b_buf, binary.LittleEndian, &x_u32)
			if err != nil {
				return result, nil
			}
			data[fields[i]] = x_u32

			result = append(result, data)
		}

		i++
	}

	return result, nil
}
