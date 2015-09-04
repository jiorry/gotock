package home

import (
	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Default struct {
	gos.Page
}

func (p *Default) RequireAuth() (string, []interface{}) {
	return "/login", nil
}

func (p *Default) Befor() bool {
	p.View.Folder = "home"
	p.Cache.Type = gos.PAGE_CACHE_FILE
	return true
}

func (p *Default) Prepare() bool {
	p.Title = "Onqee"
	// p.View.Folder = "home"
	common.SetupPage(&p.Page, "default")

	// p.Layout.TopRenderList = nil
	// p.Layout.BottomRenderList = nil
	// p.AddHead("<base href=\"/\">")
	return true
}
