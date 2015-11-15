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
	log.Level = log.LOG_ERR

	lastDate := "2015-09-30"
	ist := db.NewInsertBuilder("fin_cwzb")
	exist := db.NewExistsBuilder("fin_cwzb")
	dataset, _ := db.NewQueryBuilder("stock").Select("id,code,ctype").Limit(0).Query()

	for i, data := range dataset {
		fmt.Println("--", i, "--", data.GetString("code"), data.GetString("ctype"))
		if exist.Where("stock_id=? and date=?", data.GetInt64("id"), lastDate).Exists() {
			fmt.Println("exists", lastDate, "return")
			continue
		}

		cwzbList, zcfzbList, lrbList, xjllbList, finPercentList, err := dfcf.FetchCwfx(data.GetString("code"), data.GetString("ctype"))
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(cwzbList) == 0 {
			continue
		}

		for _, row := range cwzbList {
			row.StockId = data.GetInt64("id")

			if exist.Table("fin_cwzb").Where("stock_id=? and date=?", row.StockId, row.Date).Exists() {
				fmt.Println("exists fin_cwzb", row.Date, "return")
				continue
			}

			fmt.Println("insert fin_cwzb", row.Date)
			ist.Table("fin_cwzb").Insert(row)
		}
		for _, row := range zcfzbList {
			row.StockId = data.GetInt64("id")

			if exist.Table("fin_zcfzb").Where("stock_id=? and date=?", row.StockId, row.Date).Exists() {
				fmt.Println("exists fin_zcfzb", row.Date, "return")
				continue
			}

			fmt.Println("insert fin_zcfzb", row.Date)
			ist.Table("fin_zcfzb").Insert(row)
		}
		for _, row := range lrbList {
			row.StockId = data.GetInt64("id")

			if exist.Table("fin_lrb").Where("stock_id=? and date=?", row.StockId, row.Date).Exists() {
				fmt.Println("exists fin_lrb", row.Date, "return")
				continue
			}

			fmt.Println("insert fin_lrb", row.Date)
			ist.Table("fin_lrb").Insert(row)
		}
		for _, row := range xjllbList {
			row.StockId = data.GetInt64("id")

			if exist.Table("fin_xjllb").Where("stock_id=? and date=?", row.StockId, row.Date).Exists() {
				fmt.Println("exists fin_xjllb", row.Date, "return")
				continue
			}

			fmt.Println("insert fin_xjllb", row.Date)
			ist.Table("fin_xjllb").Insert(row)
		}
		for _, row := range finPercentList {
			row.StockId = data.GetInt64("id")

			if exist.Table("fin_percent").Where("stock_id=? and date=?", row.StockId, row.Date).Exists() {
				fmt.Println("exists fin_percent", row.Date, "return")
				continue
			}

			fmt.Println("insert fin_percent", row.Date)
			ist.Table("fin_percent").Insert(row)
		}

		if len(cwzbList) > 0 {
			time.Sleep(1 * time.Second)
		}
	}

}
