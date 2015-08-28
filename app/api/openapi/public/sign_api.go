package public

import (
	"fmt"

	"../../../lib/auth"
	"github.com/kere/gos"
	"github.com/kere/gos/lib/util"
)

type SignApi struct {
	gos.WebApi
}

func (a *SignApi) IsSecurity() bool {
	return false
}

func (a *SignApi) UserLogin(args util.MapData) (bool, error) {
	au := a.GetUserAuth().(*auth.UserAuth)

	err := au.Login(args.GetString("cipher"))
	if err != nil {
		return false, err
	}

	au.SetCookie(30 * 24 * 3600)

	return true, err
}

func (a *SignApi) Regist(args util.MapData) (bool, error) {
	appConf := gos.Configuration.GetConf("app")
	if appConf.Get("allow_regist") == "1" {
		err := auth.Regist(args.GetString("cipher"))
		return err == nil, err
	}

	return false, fmt.Errorf("管理员已经关闭了用户注册")
}
