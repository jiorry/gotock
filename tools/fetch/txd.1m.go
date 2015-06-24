package main

import (
	"../../app/lib/tools/txd"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/conf"
	"github.com/kere/gos/lib/log"
	_ "github.com/lib/pq"
)

func main() {
	c := conf.Load("../../app/app.conf")
	db.Init("app", c.GetConf("db").MapData())
	db.Current().Log = log.NewEmpty()

	txd.ScanAndStore("sh", `C:\Program360\new_tdx\vipdoc\sh\minline`, ".lc1")

	txd.ScanAndStore("sz", `C:\Program360\new_tdx\vipdoc\sz\minline`, ".lc1")
}
