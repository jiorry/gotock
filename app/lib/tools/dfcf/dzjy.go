package dfcf

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos"
	"github.com/kere/gos/db"

	iconv "github.com/djimenez/iconv-go"
)

func FillDzjy(date time.Time) error {
	l, err := FetchDzjy(date)
	if err != nil {
		return err
	}
	for _, vo := range l {
		vo.Init(vo)
		vo.Create()
		// fmt.Println(vo)
	}
	return nil
}

// FetchDzjy 抓取数据
func FetchDzjy(date time.Time) ([]*dzjyVO, error) {
	formt := "http://data.eastmoney.com/dzjy/%s.html"

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
	buy := ""
	sell := ""
	var priceNow float64
	var price float64
	var amount float64
	var total float64
	var length int

	doc.Find("#content div.list").Eq(2).Find("table tr.list_eve").Each(func(i int, tr *goquery.Selection) {
		td = tr.Find("td")
		length = td.Length()

		if length == 10 {
			dateStr = td.Eq(0).Text()
			stockCode = td.Eq(1).Text()
		} else if length == 9 {
			stockCode = td.Eq(0).Text()
		}

		row, _ = query.Where("code=?", stockCode).QueryOne()
		if row.Empty() {
			return
		}

		switch length {
		case 10:
			priceNow = util.ParseMoney(td.Eq(4).Text())
			price = util.ParseMoney(td.Eq(5).Text())
			amount = util.ParseMoney(td.Eq(6).Text())
			total = util.ParseMoney(td.Eq(7).Text())
			buy, err = iconv.ConvertString(td.Eq(8).Text(), "gb2312", "utf-8")
			if err != nil {
				return
			}
			sell, err = iconv.ConvertString(td.Eq(9).Text(), "gb2312", "utf-8")
			if err != nil {
				return
			}

		case 9:
			priceNow = util.ParseMoney(td.Eq(3).Text())
			price = util.ParseMoney(td.Eq(4).Text())
			amount = util.ParseMoney(td.Eq(5).Text())
			total = util.ParseMoney(td.Eq(6).Text())
			buy, err = iconv.ConvertString(td.Eq(7).Text(), "gb2312", "utf-8")
			if err != nil {
				return
			}
			sell, err = iconv.ConvertString(td.Eq(8).Text(), "gb2312", "utf-8")
			if err != nil {
				return
			}
		case 5:
			price = util.ParseMoney(td.Eq(0).Text())
			amount = util.ParseMoney(td.Eq(1).Text())
			total = util.ParseMoney(td.Eq(2).Text())
			buy, err = iconv.ConvertString(td.Eq(3).Text(), "gb2312", "utf-8")
			if err != nil {
				return
			}
			sell, err = iconv.ConvertString(td.Eq(4).Text(), "gb2312", "utf-8")
			if err != nil {
				return
			}
		default:
			return
		}

		dzjy = &dzjyVO{
			StockID:  row.GetInt64("id"),
			Date:     dateStr,
			PriceNow: priceNow,
			Price:    price,
			Amount:   amount,
			Total:    total,
			Buy:      buy,
			Sell:     sell,
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
	Amount   float64 `json:"amount"`    // 成交数量
	Total    float64 `json:"total"`     // 成交金额
	Buy      string  `json:"buy"`       // 买方营业部
	Sell     string  `json:"sell"`      // 卖方营业部
	StockID  int64   `json:"stock_id" skip:"update"`
}

func (a *dzjyVO) Table() string {
	return "dzjy"
}
