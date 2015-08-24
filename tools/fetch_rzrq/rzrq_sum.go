package main

import (
	"fmt"

	"../../app/lib/tools/dfcf"
)

func main() {

	result, err := dfcf.FetchRzrqSumData()
	fmt.Println(result, err)
}
