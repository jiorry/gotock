package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/cninfo"
)

func main() {
	_, err := cninfo.FetchCwzb("300003", "sz")

	if err != nil {
		fmt.Println(err)
		return
	}
}
