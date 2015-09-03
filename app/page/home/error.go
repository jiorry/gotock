package home

import (
	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Error struct {
	gos.Page
}

func (p *Error) RequireAuth() (string, []interface{}) {
	return "/login", nil
}

func (p *Error) Befor() bool {
	p.View.Folder = "home"
	p.Cache.Type = gos.PAGE_CACHE_FILE
	return true
}

func (p *Error) Prepare() bool {
	p.Title = "Stock"

	common.SetupPage(&p.Page, "default")

	p.Layout.TopRenderList = nil
	p.Layout.BottomRenderList = nil

	p.AddHead("<base href=\"/\">")

	return true
}
