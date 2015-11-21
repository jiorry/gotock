package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/hexun"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(hexun.FetchStockData("600718"))
}
