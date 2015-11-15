package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/cninfo"
	"github.com/kere/gos"
)

func main() {
	data, err := cninfo.FetchFinancialData("600718", "2015-09-30")

	if err != nil {
		gos.DoError(err).LogError()
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
