package txd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/kere/gos/db"
	"math"
	"os"
)

type Store1m struct {
	Store5m
}

func NewStore1m(prefix, folder string) *Store1m {
	s := &Store1m{}
	s.StoreBase.table = "stock_1m"
	s.StoreBase.prefix = prefix
	s.StoreBase.ext = ".lc1"
	s.StoreBase.folder = folder
	s.parent = s
	return s
}

func (a *Store1m) ReadData(name string) (db.DataSet, error) {
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
			month := math.Floor(math.Mod(float64(x_u16), 2048) / 100)
			day := math.Mod(math.Mod(float64(x_u16), 2048), 100.00)

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
