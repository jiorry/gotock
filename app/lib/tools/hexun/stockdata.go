package hexun

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/kere/gos"
	"github.com/kere/gos/db"

	iconv "github.com/djimenez/iconv-go"
)

// FillStockData 抓取数据
func FillStockData() {
	// http://stockData.stock.hexun.com/600028.shtml
	upd := db.NewUpdateBuilder("stock")
	ist := db.NewInsertBuilder("")
	q := db.NewQueryBuilder("")
	dataset, _ := db.NewQueryBuilder("stock").Select("id,code,ctype").Limit(0).Query()

	for _, data := range dataset {
		if data.GetInt("industry_id") == 0 || data.GetInt("icb_id") == 0 {
			industry, icb, err := FetchStockData(data.GetString("code"))
			if err != nil {
				gos.DoError(err)
				continue
			}

			if data.GetInt("industry_id") == 0 {
				ist.Table("industry").Insert(industry)
				row, err := q.Table("industry").Where("name=?", industry.Name).QueryOne()
				if err != nil {
					gos.DoError(err)
				} else if !row.Empty() {
					upd.Where("id=?", data.GetInt64("id")).Update(db.DataRow{"industry_id": row.GetInt64("id")})
				}
			}

			if data.GetInt("icb_id") == 0 {
				ist.Table("icb").Insert(icb)
				row, err := q.Table("icb").Where("name=?", icb.Name).QueryOne()
				if err != nil {
					gos.DoError(err)
				} else if !row.Empty() {
					upd.Where("id=?", data.GetInt64("id")).Update(db.DataRow{"icb_id": row.GetInt64("id")})
				}
			}
		}

	}
}

// FetchStockData 抓取数据
func FetchStockData(code string) (*StockIndustry, *StockICB, error) {
	// http://stockData.stock.hexun.com/600028.shtml
	formt := "http://stockData.stock.hexun.com/%s.shtml"
	resp, err := wget.Get(fmt.Sprintf(formt, code))
	if err != nil {
		return nil, nil, gos.DoError(err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, nil, gos.DoError(err)
	}

	tr := doc.Find("#list3 table.box6 tr")

	stockIndustry := &StockIndustry{}
	stockICB := &StockICB{}

	stockIndustry.Name, err = iconv.ConvertString(tr.Eq(7).Find("td").Eq(1).Text(), "gb2312", "utf-8")
	if err != nil {
		return nil, nil, gos.DoError(err)
	}

	stockICB.Name, err = iconv.ConvertString(tr.Eq(8).Find("td").Eq(1).Text(), "gb2312", "utf-8")
	if err != nil {
		return nil, nil, gos.DoError(err)
	}

	return stockIndustry, stockICB, nil
}

type StockIndustry struct {
	ID   int64  `json:"id" skip:"all"`
	Name string `json:"name"`
}
type StockICB struct {
	ID   int64  `json:"id" skip:"all"`
	Name string `json:"name"`
}
