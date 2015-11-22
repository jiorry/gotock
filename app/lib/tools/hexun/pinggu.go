package hexun

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos"
	"github.com/kere/gos/db"
)

// FillStockData 抓取数据
func FillPingGuData() {
	// http://stockData.stock.hexun.com/600028.shtml
	var stockID int64
	var pinggu *PingGuStruct
	var err error

	ist := db.NewInsertBuilder("pinggu")
	exists := db.NewExistsBuilder("pinggu")

	dataset, _ := db.NewQueryBuilder("stock").Select("id,code,ctype").Limit(0).Query()

	for _, data := range dataset {
		if exists.Where("stock_id=? and date=?", data.GetInt64("id"), util.StockQuarter()).Exists() {
			continue
		}
		pinggu, err = FetchPingGuData(data.GetString("code"))
		if err != nil {
			gos.DoError(err)
			continue
		}
		stockID = data.GetInt64("id")
		for _, row := range pinggu.Content {
			row.StockID = stockID
			ist.Insert(row)
		}
		fmt.Println("insert ", data.GetString("code"))
		time.Sleep(1 * time.Second)
	}
}

// FetchPingGuData 抓取行业评估数据
func FetchPingGuData(code string) (*PingGuStruct, error) {
	// http://stockData.stock.hexun.com/600028.shtml
	now := gos.NowInLocation()
	formt := "http://pinggu.stock.hexun.com/DataProvider/StockFinanceNewFlash.ashx?&code=%s&s=%d"
	body, err := wget.GetBody(fmt.Sprintf(formt, code, now.Unix()))
	if err != nil {
		return nil, gos.DoError(err)
	}

	v := &PingGuStruct{}
	err = xml.Unmarshal(body, v)

	if err != nil {
		return nil, gos.DoError(err)
	}

	quarter := util.StockQuarter()
	for _, item := range v.Content {
		item.Date = quarter
		switch item.Name {
		case "综合能力":
			item.Itype = 1
		case "盈利能力":
			item.Itype = 2
		case "偿债能力":
			item.Itype = 3
		case "成长能力":
			item.Itype = 4
		case "资产经营":
			item.Itype = 5
		case "市场表现":
			item.Itype = 6
		case "投资收益":
			item.Itype = 7
		}
	}

	return v, nil
}

type PingGuStruct struct {
	Title   *TitleNode    `xml:"Target>Title"`
	Content []*PingGuItem `xml:"Target>Content>Item"`
}
type TitleNode struct {
	Name string `xml:"name,attr"`
	Code string `xml:"icode,attr"`
}
type PingGuItem struct {
	Date        string  `json:"date"`
	Name        string  `xml:"name,attr" skip:"all"`
	Itype       int     `json:"itype"`
	StockID     int64   `json:"stock_id"`
	StockValue  float64 `xml:"stockvalue,attr" json:"stockvalue"`
	ICBValue    float64 `xml:"icbvalue,attr" json:"icbvalue"`
	StdValue    float64 `xml:"stdvalue,attr" json:"stdvalue"`
	AreaValue   float64 `xml:"areavalue,attr" json:"areavalue"`
	MarketValue float64 `xml:"marketvalue,attr" json:"marketvalue"`
}
