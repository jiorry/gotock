package rzrq

import (
	"fmt"

	"github.com/jiorry/gotock/app/page/common"

	"github.com/kere/gos"
)

type Stock struct {
	gos.Page
}

func (p *Stock) RequireAuth() (string, []interface{}) {
	return "/login", nil
}

// func (p *Stock) Befor() bool {
// 	p.View.Folder = "rzrq"
// 	p.Cache.Type = gos.PAGE_CACHE_FILE
// 	return true
// }

func (p *Stock) Prepare() bool {
	p.View.Folder = "rzrq"
	code := p.Ctx.RouterParam("code")
	if code == "" {
		return false
	}

	p.Title = fmt.Sprint("融资融券信息-", code)
	common.SetupPage(&p.Page, "default")
	p.AddCss(&gos.ThemeItem{Value: "jquery.jqplot.min.css"})

	return true
}
