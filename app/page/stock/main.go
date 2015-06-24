package stock

import (
	"../common"
	"github.com/kere/gos"
)

type Main struct {
	gos.Page
}

func (p *Main) Prepare() bool {
	p.Title = "Stock"
	p.View.Folder = "stock"

	common.SetupPage(&p.Page, "default")

	// if p.GetUserAuth().NotOkAndRedirect("/error/a") {
	// 	return false
	// }

	p.Layout.TopRenderList = nil
	// p.Layout.BottomRenderList = nil

	p.AddHead("<base href=\"/\">")

	return true
}
