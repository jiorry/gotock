package auth

import (
	"github.com/kere/gos"
	"github.com/kere/gos/db"
	"github.com/kere/gos/lib/util"
	"time"
)

func QueryUserById(id int64) db.DataRow {
	return New(nil).QueryById(id)
}

func QueryUserByNick(n string) db.DataRow {
	return New(nil).Query(n)
}

func QueryByEmail(email string) db.DataRow {
	return New(nil).QueryByEmail(email)
}

type UserVO struct {
	db.BaseVO
	Id       int64     `json:"id"`
	Nick     string    `json:"nick"`
	Avatar   string    `json:"avatar"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Status   int8      `json:"status"`
	JsonData string    `json:"json_data"`
	Created  time.Time `json:"created_at"`
	LastSee  time.Time `json:"last_see_at"`
}

func NewUserVO(u db.DataRow) *UserVO {
	if u.Empty() {
		return &UserVO{
			Id:     int64(0),
			Nick:   "empty",
			Avatar: "empty",
			Email:  "empty@xxx.xxx",
		}
	}
	vo := &UserVO{}
	u.CopyToStruct(vo)
	return vo
}

type UserAuth struct {
	gos.UserAuth
}

func UserToken(email string, text []byte) (string, string) {
	salt := util.Unique()
	return salt, gos.UserToken([]interface{}{email}, text, []byte(salt))
}

func New(ctx *gos.Context) *UserAuth {
	auth := (&UserAuth{})
	auth.Init(ctx)
	vo := auth.GetOptions()
	auth.SetKeys(vo.FieldNick)
	return auth
}

// ---------------------------------
func (u *UserAuth) GetUserVO() *UserVO {
	return NewUserVO(u.User())
}

func (u *UserAuth) SetCookie(age int64) {
	u.UserAuth.UserAuthBase.SetCookie(u.Keys(), u.User(), age)
}

func (u *UserAuth) Avatar() string {
	return "http://localhost:8080/assets/uploads/avatar/face.png"
}

func (u *UserAuth) Nick() string {
	return u.User().GetString("nick")
}

func (a *UserAuth) Login(cipher string) error {
	return a.UserAuth.UserAuthBase.LoginBy([]string{"email"}, []string{"email"}, []byte(cipher))
}

func (a *UserAuth) QueryById(id int64) db.DataRow {
	r, _ := db.NewQueryBuilder(a.GetOptions().Table).Where("id=?", id).CacheExpire(a.GetExpire()).QueryOne()
	return r
}

func (a *UserAuth) Query(nick string) db.DataRow {
	return a.QueryByKeys([]string{"nick"}, []interface{}{nick})
}

func (a *UserAuth) QueryByEmail(email string) db.DataRow {
	return a.QueryByKeys([]string{"email"}, []interface{}{email})
}

func (a *UserAuth) ClearCache() {
	a.ClearCacheByKeys([]string{"nick"}, []interface{}{a.User().GetString("nick")})
	a.ClearCacheByKeys([]string{"email"}, []interface{}{a.User().GetString("email")})
	db.NewQueryBuilder(a.GetOptions().Table).Where("id=?", a.UserId()).ClearCache()
}

func (a *UserAuth) UserData() db.DataRow {
	if a.User().IsSet("branches") {
		return a.User()
	} else {
		a.SetCurrentUser(a.QueryById(a.User().GetInt64("id")))
	}
	return a.User()
}

func (a *UserAuth) CookieLang() string {
	cookie, err := a.GetContext().Request.Cookie("lang")
	if err != nil {
		return "en-US"
	}
	return cookie.Value
}
