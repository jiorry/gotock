package main

import (
	"../../app/lib/tools/txd"
	"fmt"
)

func main() {
	folder := `C:\Program360\new_tdx\vipdoc\sh\lday`

	result, err := txd.ScanDirGetAllFileName(folder, ".day")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(len(result), " records to write")
	for _, item := range result {
		fmt.Println(item)
	}
}
