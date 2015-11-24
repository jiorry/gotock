package dfcf

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos"
	"github.com/kere/gos/db"
)

func FillDzjy(date time.Time) error {
	l, err := FetchDzjy(date)
	if err != nil {
		return err
	}
	for _, vo := range l {
		vo.Init(vo)
		vo.Create()
	}
	return nil
}

// FetchDzjy 抓取数据
func FetchDzjy(date time.Time) ([]*dzjyVO, error) {
	formt := "http://data.eastmoney.com/dzjy/%d.html"

	resp, err := wget.Get(fmt.Sprintf(formt, date.Format("200601")))
	if err != nil {
		return nil, gos.DoError(err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, gos.DoError(err)
	}
	var td *goquery.Selection
	var dzjy *dzjyVO
	var dateStr string
	var stockCode string
	var row db.DataRow
	query := db.NewQueryBuilder("stock")
	datalist := make([]*dzjyVO, 0)

	doc.Find("#content div.list").Eq(2).Find("table tr.list_eve").Each(func(i int, tr *goquery.Selection) {
		td = tr.Find("td")
		if td.Length() == 10 {
			dateStr = td.Eq(0).Text()
			stockCode = td.Eq(1).Text()
		}

		row, _ = query.Where("code=?", stockCode).QueryOne()
		if row.Empty() {
			return
		}

		dzjy = &dzjyVO{
			StockID:  row.GetInt64("id"),
			Date:     dateStr,
			PriceNow: util.ParseMoney(td.Eq(4).Text()),
			Price:    util.ParseMoney(td.Eq(5).Text()),
			Amount:   util.ParseMoney(td.Eq(6).Text()),
			Total:    util.ParseMoney(td.Eq(7).Text()),
			Buy:      td.Eq(8).Text(),
			Sell:     td.Eq(9).Text(),
		}

		datalist = append(datalist, dzjy)
	})

	return datalist, nil
}

type dzjyVO struct {
	db.BaseVO
	Date     string  `json:"date"`
	PriceNow float64 `json:"price_now"` // 现价
	Price    float64 `json:"price"`     // 成交价格
	Amount   float64 `json:"price_now"` // 成交数量
	Total    float64 `json:"price_now"` // 成交金额
	Buy      string  `json:"buy"`       // 买方营业部
	Sell     string  `json:"sell"`      // 卖方营业部
	StockID  int64   `json:"stock_id" skip:"update"`
}

func (a *dzjyVO) Table() string {
	return "dzjy"
}
