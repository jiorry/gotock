package user

import (
	"../common"

	"github.com/kere/gos"
)

type Login struct {
	gos.Page
}

func (p *Login) Prepare() bool {
	p.Title = "用户登录"
	p.View.Folder = "user"
	common.SetupPage(&p.Page, "default")

	return true
}
