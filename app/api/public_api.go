package api

import (
	"fmt"
	"github.com/kere/gos"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/util"
	"time"
)

type Public struct {
	gos.WebApi
}

func (a *Public) IsExists(args util.MapData) (int, error) {
	table := "users"
	field := args.GetString("field")
	val := args.GetString("value")

	exists := db.NewExistsBuilder(table).Where(field+"=?", val)

	if isEx := exists.Exists(); isEx {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (a *Public) Rsakey() (map[string]interface{}, error) {
	k := gos.GetRSAKey(0)
	m := make(map[string]interface{}, 0)
	m["hex"] = fmt.Sprintf("%x", k.Key.PublicKey.N)
	m["keyid"] = k.CreatedAt.Unix()
	m["unix"] = time.Now().Unix()
	return m, nil
}
