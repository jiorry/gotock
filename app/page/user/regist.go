package user

import (
	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Regist struct {
	gos.Page
}

func (p *Regist) Prepare() bool {
	p.Title = "用户注册"
	p.View.Folder = "user"
	common.SetupPage(&p.Page, "default")

	return true
}
