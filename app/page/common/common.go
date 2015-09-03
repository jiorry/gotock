package common

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jiorry/gotock/app/lib/auth"

	"github.com/kere/gos"
	"github.com/kere/gos/lib/util"
)

var (
	b_s1 = []byte("<script>var MYENV='")
	b_s2 = []byte("',THEME='")
	b_s3 = []byte("',JsVersion='")
	b_s4 = []byte("',Server = {diff : 0,setTime : function(unix){this.diff = unix>0 ? (new Date()).getTime() - unix*1000 : 0;},getTime : function(time){return (new Date()).getTime() - this.diff;}};")
	b_s5 = []byte("</script>\n")

	b_d1 = []byte("<div id=\"wrap\">\n")
	b_d2 = []byte("</div>\n")

	rq_a1 = []byte("<script type=\"text/javascript\">var require={urlArgs:\"")
	rq_a2 = []byte("\"};</script>\n")
	rq_b1 = []byte("<script src=\"/assets/js/require.js\" data-main=\"")
	rq_b2 = []byte("\" defer=\"\" async=\"true\"></script>\n")
)

func SetupPage(p *gos.Page, theme string) {
	if theme == "" {
		theme = gos.GetSite().SiteTheme
	}
	p.View.Theme = theme
	p.JsPosition = "end"

	// p.AddHead("<link href=\"//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css\" rel=\"stylesheet\">")
	p.AddCss(&gos.ThemeItem{Value: "bootstrap.min.css"})
	p.AddCss(&gos.ThemeItem{Value: "font-awesome.min.css"})

	p.AddCss(&gos.ThemeItem{Value: fmt.Sprint(p.View.Value, ".css"), Folder: fmt.Sprint(gos.RunMode, "/page/", p.View.Folder), Theme: theme})

	p.Layout.AddTopRender(gos.NewTemplateRender("", "", "_header", nil))
	// p.Layout.AddBottomRender(gos.NewTemplateRender("", "", "_footer", nil))

	p.Layout.RenderFunc = Render
	p.SetUserAuth(auth.New(p.Ctx))

	RequireJs(p)
	p.AddJs(&gos.ThemeItem{Value: "jquery.js"})
}

func UserAuth(ctx *gos.Context) *auth.UserAuth {
	return auth.New(ctx)
}

func ToError(page *gos.Page, msg string) {
	page.Ctx.Redirect("/error?message=%s", msg)
}

func RequireJs(p *gos.Page) {
	url := ""
	if gos.RunMode == "dev" {
		url = util.UrlJoin(gos.GetSite().StaticUrl, "/assets/js/dev/page/", p.View.Folder, p.View.Value)
	} else {
		url = util.UrlJoin(gos.GetSite().StaticUrl, "/assets/js/pro/page/", p.View.Folder, p.View.Value)
	}

	ver := gos.GetSite().JsVersion
	s := bytes.Buffer{}
	if ver != "" {
		s.Write(rq_a1)
		s.WriteString("v=")
		s.WriteString(ver)
		s.Write(rq_a2)
	}

	s.Write(rq_b1)
	s.WriteString(url)
	s.Write(rq_b2)

	textRender := &gos.TextRender{
		Name:   "requirejs",
		Source: string(s.Bytes()),
		Data:   nil,
	}

	p.Layout.AddBottomRender(textRender)
}

func Render(p *gos.Page, w io.Writer) {
	a := p.Layout

	w.Write(gos.B_HTML_BEGIN)
	a.HeadLayout.Render(w)
	w.Write(gos.B_HTML_BODY_BEGIN)

	if len(a.TopRenderList) > 0 {
		for _, r := range a.TopRenderList {
			r.Render(w)
		}
	}

	w.Write(b_d1)
	if len(a.ContentRenderList) > 0 {
		for _, r := range a.ContentRenderList {
			r.Render(w)
		}
	}
	w.Write(b_d2)

	if len(a.BottomRenderList) > 0 {
		for _, r := range a.BottomRenderList {
			r.Render(w)
		}
	}

	w.Write(b_s1)
	w.Write([]byte(gos.RunMode))
	w.Write(b_s2)
	w.Write([]byte(gos.GetSite().SiteTheme))
	w.Write(b_s3)
	w.Write([]byte(gos.GetSite().JsVersion))
	w.Write(b_s4)
	w.Write(b_s5)

	a.HeadLayout.RenderBottomJs(w)
	w.Write(gos.B_HTML_BODY_END)
	w.Write(gos.B_HTML_END)
}
