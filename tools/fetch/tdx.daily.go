package main

import (
	"github.com/jiorry/gotock/app/lib/tools/tdx"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/conf"
	"github.com/kere/gos/lib/log"
	_ "github.com/lib/pq"
)

func main() {
	c := conf.Load("../../app/app.conf")
	db.Init("app", c.GetConf("db").MapData())
	db.Current().Log = log.NewEmpty()

	tdx.ScanAndStore("sh", `C:\Program360\new_tdx\vipdoc\sh\lday`, ".day")

	tdx.ScanAndStore("sz", `C:\Program360\new_tdx\vipdoc\sz\lday`, ".day")
}
