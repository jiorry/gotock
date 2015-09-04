package user

import (
	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Login struct {
	gos.Page
}

// func (p *Login) Befor() bool {
// 	p.View.Folder = "user"
// 	p.Cache.Type = gos.PAGE_CACHE_FILE
// 	return true
// }

func (p *Login) Prepare() bool {
	p.Title = "用户登录"
	p.View.Folder = "user"
	common.SetupPage(&p.Page, "default")

	p.Layout.TopRenderList = nil

	return true
}
