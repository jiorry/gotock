package auth

import (
	"bytes"
	"time"

	"github.com/kere/gos"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/util"
)

func QueryUserById(id int64) db.DataRow {
	return New(nil).QueryById(id)
}

type UserAuth struct {
	gos.UserAuth
}

func New(ctx *gos.Context) *UserAuth {
	auth := (&UserAuth{})
	auth.Init(ctx)
	vo := auth.GetOptions()
	auth.SetKeys(vo.FieldNick)
	return auth
}

// ---------------------------------
// func (u *UserAuth) SetCookie(age int64) {
// 	u.UserAuth.UserAuthBase.SetCookie(u.Keys(), u.User(), age)
// }

func (a *UserAuth) Login(cipher string) error {
	return a.UserAuth.LoginBy([]string{"nick"}, []string{"nick"}, []byte(cipher))
}

func (a *UserAuth) QueryById(id int64) db.DataRow {
	r, _ := db.NewQueryBuilder(a.GetOptions().Table).Where("id=?", id).CacheExpire(a.GetExpire()).QueryOne()
	return r
}

func (a *UserAuth) Query(nick string) db.DataRow {
	return a.QueryByKeys([]string{"nick"}, []interface{}{nick})
}

func (a *UserAuth) ClearCache() {
	a.ClearCacheByKeys([]string{"nick"}, []interface{}{a.User().GetString("nick")})
	db.NewQueryBuilder(a.GetOptions().Table).Where("id=?", a.UserId()).ClearCache()
}

// func (a *UserAuth) CookieLang() string {
// 	cookie, err := a.GetContext().Request.Cookie("lang")
// 	if err != nil {
// 		return "en-US"
// 	}
// 	return cookie.Value
// }
var separator = []byte("|")

func Regist(cipher string) error {
	_, b, err := gos.PraseCipher([]byte(cipher))
	if err != nil {
		return gos.NewError(0, err)
	}

	arr := bytes.Split(b, separator)

	obj := db.DataRow{}
	obj["nick"] = string(arr[0])
	obj["status"] = 1
	obj["created_at"] = time.Now()
	obj["salt"], obj["token"] = UserToken(string(arr[0]), arr[1])
	_, err = db.NewInsertBuilder("users").Insert(obj)
	return err
}

func UserToken(nick string, pwd []byte) (string, string) {
	salt := util.Unique()
	return salt, gos.UserToken([]interface{}{nick}, pwd, []byte(salt))
}
