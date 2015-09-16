package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/dfcf"
)

func main() {
	items, err := dfcf.GetHgtAmount()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(items) == 0 {
		return
	}
	fmt.Println(len(items), items[0])

}
