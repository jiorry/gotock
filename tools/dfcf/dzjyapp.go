package main

import (
	"fmt"

	"github.com/jiorry/gotock/app/lib/tools/dfcf"
	"github.com/kere/gos"
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

	fmt.Println(dfcf.FillDzjy(gos.NowInLocation()), "finished")

}
