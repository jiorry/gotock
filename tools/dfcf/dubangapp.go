package main

import (
	"fmt"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/dfcf"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/conf"
	"github.com/kere/gos/lib/log"
	_ "github.com/lib/pq"
)

func main() {
	c := conf.Load("../../app/app.conf")
	db.Init("app", c.GetConf("db").MapData())
	db.Current().Log = log.NewEmpty()

	lastDate := "2015-09-30"
	ist := db.NewInsertBuilder("fin_dubang")
	exist := db.NewExistsBuilder("fin_dubang")
	dataset, _ := db.NewQueryBuilder("stock").Select("id,code,ctype").Limit(0).Query()
	isIst := false

	for i, data := range dataset {
		fmt.Println("--", i, "--", data.GetString("code"))
		if exist.Where("stock_id=? and date=?", data.GetInt64("id"), lastDate).Exists() {
			fmt.Println("exists", lastDate, "return")
			continue
		}

		dubangList, err := dfcf.FetchDubang(data.GetString("code"), data.GetString("ctype"))
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(dubangList) == 0 {
			continue
		}

		isIst = false
		for _, dubang := range dubangList {
			dubang.StockId = data.GetInt64("id")

			if exist.Where("stock_id=? and date=?", dubang.StockId, dubang.Date).Exists() {
				fmt.Println("exists", dubang.Date, "return")
				continue
			}

			fmt.Println("insert", dubang.Date)
			ist.Insert(dubang)
			isIst = true
		}

		if isIst {
			time.Sleep(1 * time.Second)
		}
	}

}
