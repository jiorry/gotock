package main

import (
	"fmt"

	"strings"

	"github.com/jiorry/gotock/app/lib/tools/dzh"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/conf"
	"github.com/kere/gos/lib/log"
	_ "github.com/lib/pq"
)

func main() {
	name := `C:\Program360\dzh2_hlzq\Download\FIN\full.FIN`
	l, err := dzh.ReadData(name)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	c := conf.Load("../../app/app.conf")
	db.Init("app", c.GetConf("db").MapData())
	db.Current().Log = log.NewEmpty()

	rows := make([]db.DataRow, 100)

	// insert := db.NewInsertBuilder("stock")
	var k int = 0

	for _, finData := range l {
		if k == 100 {
			// insert.InsertM(rows)
			fmt.Println(finData)
			k = 0
			rows = make([]db.DataRow, 100)
		}

		if finData.Fin3 == 0 {
			continue
		}

		rows[k] = db.DataRow{"code": finData.Code[2:], "ctype": strings.ToLower(finData.Code[:2])}
		k++
	}

	if k > 0 {
		fmt.Println(k, "finished")
		// insert.InsertM(rows)
	}
}
