package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/tdx"
)

func main() {
	folder := `C:\Program360\new_tdx\vipdoc\sh\lday`

	result, err := tdx.ScanDirGetAllFileName(folder, ".day")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(len(result), " records to write")
	for _, item := range result {
		fmt.Println(item)
	}
}
