package api

import (
	"../lib/auth"
	"./openapi/index"
	"fmt"
	"github.com/kere/gos"
	"reflect"
)

var apiMap = make(map[string]reflect.Type)

func regist(n string, a interface{}) {
	typ := reflect.TypeOf(a)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	apiMap[n] = typ
}

func init() {
	regist("index.SZApi", &index.SZApi{})
	regist("index.SHApi", &index.SHApi{})
}

type OpenApi struct {
	gos.OpenApi
}

func (a *OpenApi) Prepare() bool {
	a.SetUserAuth(auth.New(a.WebApi.Ctx))
	return true
}

func (a *OpenApi) Factory(n string) (gos.IApi, error) {
	if v, ok := apiMap[n]; ok {
		api := reflect.New(v).Interface().(gos.IApi)

		if api.IsSecurity() && a.GetUserAuth().NotOk() {
			return nil, fmt.Errorf("用户没有登录或登录已过期，不能使用此API")
		}

		api.SetUserAuth(a.GetUserAuth())
		return api, nil
	}
	return nil, fmt.Errorf("api %s not registed", n)
}
