package home

import (
	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Default struct {
	gos.Page
}

func (p *Default) Prepare() bool {
	p.Title = "Stock"
	p.View.Folder = "home"

	common.SetupPage(&p.Page, "default")

	// p.Layout.TopRenderList = nil
	p.Layout.BottomRenderList = nil
	// p.AddHead("<base href=\"/\">")

	if p.GetUserAuth().NotOk() {
		p.Ctx.Redirect("/login")
		return false
	}

	return true
}
