package dfcf

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
)

func init() {

}

// FetchRzrqSumData 抓取数据
func FetchDubang(code, ctype string) ([]*Dubang, error) {
	c := "01"
	if strings.ToLower(ctype) == "sz" {
		c = "02"
	}
	formt := "http://soft-f9.eastmoney.com/soft/gp61.php?code=%s%s"

	u := fmt.Sprintf(formt, code, c)

	resp, err := wget.Get(u)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	dataset := make([]*Dubang, 0)

	doc.Find("#tablefont tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			s.Find("td").Each(func(k int, td *goquery.Selection) {
				if k == 0 {
					return
				}

				dataset = append(dataset, &Dubang{Date: fmt.Sprint("20", td.Find("p span").Text())})
			})

			return
		}

		s.Find("td").Each(func(k int, td *goquery.Selection) {
			if k == 0 {
				return
			}

			reflect.ValueOf(dataset[k-1]).
				Elem().Field(i).SetFloat(util.ParseMoneyCN(td.Find("p span").Text()))
		})
	})

	return dataset, nil
}

type Dubang struct {
	Date string  `json:"date"`
	A1   float64 `json:"a1"`  // 净资产收益率 %
	A2   float64 `json:"a2"`  // 总资产净利率 %
	A3   float64 `json:"a3"`  // 营业净利润率(%)
	A4   float64 `json:"a4"`  // 净利润(元)
	A5   float64 `json:"a5"`  // 收入总额(元)
	A6   float64 `json:"a6"`  // 营业收入(元)
	A7   float64 `json:"a7"`  // 公允价值变动收益(元)
	A8   float64 `json:"a8"`  // 营业外收入(元)
	A9   float64 `json:"a9"`  // 投资收益(元)
	A10  float64 `json:"a10"` // 成本总额(元)
	A11  float64 `json:"a11"` // 营业成本(元)
	A12  float64 `json:"a12"` // 营业税金及附加(元)
	A13  float64 `json:"a13"` // 所得税费用(元)
	A14  float64 `json:"a14"` // 资产减值损失(元)
	A15  float64 `json:"a15"` // 营业外支出(元)
	A16  float64 `json:"a16"` // 期间费用(元)
	A17  float64 `json:"a17"` // 财务费用(元)
	A18  float64 `json:"a18"` // 销售费用(元)
	A19  float64 `json:"a19"` // 管理费用(元)
	A20  float64 `json:"a20"` // 营业收入(元)
	A21  float64 `json:"a21"` // 总资产周转率(%)
	A22  float64 `json:"a22"` // 营业收入(元)
	A23  float64 `json:"a23"` // 资产总额(元)
	A24  float64 `json:"a24"` // 流动资产(元)
	A25  float64 `json:"a25"` // 货币资金(元)
	A26  float64 `json:"a26"` // 交易性金融资产(元)
	A27  float64 `json:"a27"` // 应收账款(元)
	A28  float64 `json:"a28"` // 预付账款(元)
	A29  float64 `json:"a29"` // 其他应收款(元)
	A30  float64 `json:"a30"` // 存货(元)
	A31  float64 `json:"a31"` // 其他流动资产(元)
	A32  float64 `json:"a32"` // 非流动资产(元)
	A33  float64 `json:"a33"` // 可供出售金融资产(元)
	A34  float64 `json:"a34"` // 持有至到期投资(元)
	A35  float64 `json:"a35"` // 长期股权投资(元)
	A36  float64 `json:"a36"` // 投资性房地产(元)
	A37  float64 `json:"a37"` // 固定资产(元)
	A38  float64 `json:"a38"` // 在建工程(元)
	A39  float64 `json:"a39"` // 无形资产(元)
	A40  float64 `json:"a40"` // 开发支出(元)
	A41  float64 `json:"a41"` // 商誉(元)
	A42  float64 `json:"a42"` // 长期待摊费用(元)
	A43  float64 `json:"a43"` // 递延所得税资产(元)
	A44  float64 `json:"a44"` // 其他非流动资产(元)
	A45  float64 `json:"a45"` // 权益乘数
	A46  float64 `json:"a46"` // 资产负债率(%)
	A47  float64 `json:"a47"` // 负债总额(元)
	A48  float64 `json:"a48"` // 资产总额(元)

	StockId int64 `json:"stock_id" skip:"update"`
}
