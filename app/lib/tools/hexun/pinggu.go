package hexun

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/kere/gos"
)

// FetchPingGuData 抓取行业评估数据
func FetchPingGuData(code string) (*PingGuStruct, error) {
	// http://stockData.stock.hexun.com/600028.shtml
	formt := "http://pinggu.stock.hexun.com/DataProvider/StockFinanceNewFlash.ashx?&code=%s&s=%d"
	body, err := wget.GetBody(fmt.Sprintf(formt, code, time.Now().Unix()))
	if err != nil {
		return nil, gos.DoError(err)
	}

	v := PingGuStruct{}
	err = xml.Unmarshal(body, &v)

	if err != nil {
		return nil, gos.DoError(err)
	}

	return &v, nil
}

type PingGuStruct struct {
	Title   TitleNode    `xml:"Target>Title"`
	Content []PingGuItem `xml:"Target>Content>Item"`
}
type TitleNode struct {
	Name string `xml:"name,attr"`
	Code string `xml:"icode,attr"`
}
type PingGuItem struct {
	Name        string  `xml:"name,attr"`
	StockValue  float64 `xml:"stockvalue,attr"`
	ICBValue    float64 `xml:"icbvalue,attr"`
	StdValue    float64 `xml:"stdvalue,attr"`
	AreaValue   float64 `xml:"areavalue,attr"`
	MarketValue float64 `xml:"marketvalue,attr"`
}
