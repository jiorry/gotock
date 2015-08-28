package home

import (
	"../common"

	"github.com/kere/gos"
)

type Error struct {
	gos.Page
}

func (p *Error) Prepare() bool {
	p.Title = "Stock"
	p.View.Folder = "home"

	common.SetupPage(&p.Page, "default")

	// if p.GetUserAuth().NotOkAndRedirect("/error/a") {
	// 	return false
	// }

	// p.Layout.TopRenderList = nil
	// p.Layout.BottomRenderList = nil

	p.AddHead("<base href=\"/\">")

	return true
}
