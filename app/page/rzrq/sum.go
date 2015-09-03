package rzrq

import (
	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Sum struct {
	gos.Page
}

func (p *Sum) RequireAuth() (string, []interface{}) {
	return "/login", nil
}

func (p *Sum) Prepare() bool {
	p.View.Folder = "rzrq"
	p.Title = "两市融资融券信息"
	common.SetupPage(&p.Page, "default")
	p.AddCss(&gos.ThemeItem{Value: "jquery.jqplot.min.css"})

	return true
}
