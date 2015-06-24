package dzh

import (
	// "bytes"
	// "encoding/binary"
	"fmt"
	"github.com/kere/gos/db"
	"os"
)

func ReadData(name string) (db.DataSet, error) {
	result := db.DataSet{}

	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return result, err
	}

	defer file.Close()

	code := make([]byte, 240)
	file.Read(code)

	// b_buf := bytes.NewBuffer(code)
	// // binary.Read(b_buf, binary.BigEndian, &x)

	// var x int32
	// err = binary.Read(b_buf, binary.LittleEndian, &x)

	fmt.Printf("%s", code)

	return result, nil
}
