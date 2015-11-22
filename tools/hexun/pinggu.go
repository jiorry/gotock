package main

import (
	"github.com/jiorry/gotock/app/lib/tools/hexun"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/conf"
	"github.com/kere/gos/lib/log"
	_ "github.com/lib/pq"
)

func main() {
	c := conf.Load("../../app/app.conf")
	db.Init("app", c.GetConf("db").MapData())
	// db.Current().Log = log.NewEmpty()
	log.Level = log.LOG_ERR

	hexun.FillPingGuData()
}
