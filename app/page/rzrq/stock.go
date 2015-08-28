package rzrq

import (
	"../common"

	"github.com/kere/gos"
)

type Stock struct {
	gos.Page
}

func (p *Stock) Prepare() bool {
	p.Title = "两市融资融券信息"
	p.View.Folder = "rzrq"
	common.SetupPage(&p.Page, "default")

	return true
}
