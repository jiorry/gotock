package cninfo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiorry/gotock/app/lib/tools/wget"
)

func init() {

}

// FetchRzrqSumData 抓取数据
func FetchCwzb(code, ctype string) ([]*Cwzb, error) {
	ctypeStr := "K"
	if ctype == "sz" {
		ctypeStr = "J"
	}
	formt := "http://quotes.cnfol.com/new/f10/cwzb/%s%s.html"

	resp, err := wget.Get(fmt.Sprintf(formt, code, ctypeStr))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() < 5 {
			return
		}

		td.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			fmt.Println(s.Text())
		})
	})

	return nil, nil
}

type Cwzb struct {
	Date string  `json:"date"`
	A1   float64 `json:"a1"`  // 每股收益(元)
	A2   float64 `json:"a2"`  // 每股收益扣除(元)
	A3   float64 `json:"a3"`  // 每股加权收益(元)
	A4   float64 `json:"a4"`  // 每股净资产(元)
	A5   float64 `json:"a5"`  // 每股资本公积金(元)
	A6   float64 `json:"a6"`  // 每股未分配利润(元)
	A7   float64 `json:"a7"`  // 每股经营活动净流量(元)
	A8   float64 `json:"a8"`  // 每股现金净流量(元)
	A9   float64 `json:"a9"`  // 营业成本率(%)
	A10  float64 `json:"a10"` // 主营业务成本占主营业务收入比例(%)
	A11  float64 `json:"a11"` // 营业利润与利润总额的比例(%)
	A12  float64 `json:"a12"` // 投资收益与利润总额的比例(%)
	A13  float64 `json:"a13"` // 补贴收入与利润总额的比例(%)
	A14  float64 `json:"a14"` // 营业外收支与利润总额的比例(%)
	A15  float64 `json:"a15"` // 净利润与利润总额的比例(%)
	A16  float64 `json:"a16"` // 销售毛利率(%)
	A17  float64 `json:"a17"` // 销售净利率(%)
	A18  float64 `json:"a18"` // 总资产收益率(%)
	A19  float64 `json:"a19"` // 净资产收益率(%)
	A20  float64 `json:"a20"` // 存货周转率(%)
	A21  float64 `json:"a21"` // 应收帐款周转率(%)
	A22  float64 `json:"a22"` // 总资产周转率(%)
	A23  float64 `json:"a23"` // 固定资产周转率(%)
	A24  float64 `json:"a24"` // 股东权益周转率(%)
	A25  float64 `json:"a25"` // 营业利润增长率(%)
	A26  float64 `json:"a26"` // 税后利润增长率(%)
	A27  float64 `json:"a27"` // 净资产增长率(%)
	A28  float64 `json:"a28"` // 利润总额增长率(%)
	A29  float64 `json:"a29"` // 总资产增长率(%)
	A30  float64 `json:"a30"` // 流动比率(%)
	A31  float64 `json:"a31"` // 速动比率(%)
	A32  float64 `json:"a32"` // 利息保障倍数
	A33  float64 `json:"a33"` // 股东权益与固定资产比率(%)
	A34  float64 `json:"a34"` // 长期负债与运营资金比例(%)

}
