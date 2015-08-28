package public

import (
	"fmt"
	"time"

	"github.com/kere/gos"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/util"
)

type PublicApi struct {
	gos.WebApi
}

func (a *PublicApi) IsSecurity() bool {
	return false
}

func (a *PublicApi) IsExists(args util.MapData) (int, error) {
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

func (a *PublicApi) Rsakey() (map[string]interface{}, error) {
	k := gos.GetRSAKey(0)
	m := make(map[string]interface{}, 0)
	m["hex"] = fmt.Sprintf("%x", k.Key.PublicKey.N)
	m["keyid"] = k.CreatedAt.Unix()
	m["unix"] = time.Now().Unix()

	isOk := a.GetUserAuth().IsOk()
	m["is_login"] = isOk

	return m, nil
}
