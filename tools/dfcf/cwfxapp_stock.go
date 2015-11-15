package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/dfcf"
)

func main() {
	// cwzbList, zcfzbList, lrbList, xjllbList, finPercentList, err := dfcf.FetchCwfx("603005", "sh")
	cwzbList, zcfzbList, lrbList, xjllbList, finPercentList, err := dfcf.FetchCwfx("000001", "sz")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(cwzbList))
	fmt.Println(len(zcfzbList))
	fmt.Println(len(lrbList))
	fmt.Println(len(xjllbList))
	fmt.Println(len(finPercentList))
}
