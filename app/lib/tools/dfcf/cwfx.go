package dfcf

import (
	"fmt"
	"reflect"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos"
	"github.com/kere/gos/db"
)

func init() {

}

// FillCwfx 抓取并且填充数据
func FillCwfx(lastDate string) {
	ist := db.NewInsertBuilder("fin_cwzb")
	exist := db.NewExistsBuilder("fin_cwzb")
	dataset, _ := db.NewQueryBuilder("stock").Select("id,code,ctype").Limit(0).Query()

	for i, data := range dataset {
		fmt.Println("--", i, "--", data.GetString("code"), data.GetString("ctype"))
		if exist.Where("stock_id=? and date=?", data.GetInt64("id"), lastDate).Exists() {
			gos.Log.Info("exists", lastDate, "return")
			continue
		}

		cwzbList, zcfzbList, lrbList, xjllbList, finPercentList, err := FetchCwfx(data.GetString("code"), data.GetString("ctype"))
		if err != nil {
			gos.DoError(err)
			continue
		}

		if len(cwzbList) == 0 {
			continue
		}

		for _, row := range cwzbList {
			row.StockID = data.GetInt64("id")

			if exist.Table("fin_cwzb").Where("stock_id=? and date=?", row.StockID, row.Date).Exists() {
				gos.Log.Info("exists fin_cwzb", row.Date, "return")
				continue
			}

			gos.Log.Info("insert fin_cwzb", row.Date)
			ist.Table("fin_cwzb").Insert(row)
		}
		for _, row := range zcfzbList {
			row.StockID = data.GetInt64("id")

			if exist.Table("fin_zcfzb").Where("stock_id=? and date=?", row.StockID, row.Date).Exists() {
				gos.Log.Info("exists fin_zcfzb", row.Date, "return")
				continue
			}

			gos.Log.Info("insert fin_zcfzb", row.Date)
			ist.Table("fin_zcfzb").Insert(row)
		}
		for _, row := range lrbList {
			row.StockID = data.GetInt64("id")

			if exist.Table("fin_lrb").Where("stock_id=? and date=?", row.StockID, row.Date).Exists() {
				fmt.Println("exists fin_lrb", row.Date, "return")
				continue
			}

			gos.Log.Info("insert fin_lrb", row.Date)
			ist.Table("fin_lrb").Insert(row)
		}
		for _, row := range xjllbList {
			row.StockID = data.GetInt64("id")

			if exist.Table("fin_xjllb").Where("stock_id=? and date=?", row.StockID, row.Date).Exists() {
				gos.Log.Info("exists fin_xjllb", row.Date, "return")
				continue
			}

			gos.Log.Info("insert fin_xjllb", row.Date)
			ist.Table("fin_xjllb").Insert(row)
		}
		for _, row := range finPercentList {
			row.StockID = data.GetInt64("id")

			if exist.Table("fin_percent").Where("stock_id=? and date=?", row.StockID, row.Date).Exists() {
				gos.Log.Info("exists fin_percent", row.Date, "return")
				continue
			}

			gos.Log.Info("insert fin_percent", row.Date)
			ist.Table("fin_percent").Insert(row)
		}

		if len(cwzbList) > 0 {
			time.Sleep(1 * time.Second)
		}
	}
}

// FetchRzrqSumData 抓取数据
func FetchCwfx(code, ctype string) ([]*Cwzb, []*Zcfzb, []*Lrb, []*Xjllb, []*FinPercent, error) {
	formt := "http://f10.eastmoney.com/f10_v2/FinanceAnalysis.aspx?code=%s%s"

	resp, err := wget.Get(fmt.Sprintf(formt, ctype, code))
	if err != nil {
		return nil, nil, nil, nil, nil, gos.DoError(err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, nil, nil, nil, nil, gos.DoError(err)
	}

	cwzbList := make([]*Cwzb, 0)
	zcfzbList := make([]*Zcfzb, 0)
	lrbList := make([]*Lrb, 0)
	xjllbList := make([]*Xjllb, 0)
	finPercentList := make([]*FinPercent, 0)
	index := 0

	doc.Find("#F10MainTargetDiv table tr").Each(func(i int, tr *goquery.Selection) {
		if i == 0 {
			tr.Find("th.tips-fieldname-Right").Each(func(k int, th *goquery.Selection) {
				cwzbList = append(cwzbList, &Cwzb{Date: fmt.Sprint("20", th.Text())})
			})
			return
		}

		if _, isOk := tr.Attr("onclick"); isOk {
			tr.Find("td.tips-data-Right").Each(func(k int, td *goquery.Selection) {
				reflect.ValueOf(cwzbList[k]).
					Elem().Field(index + 1).SetFloat(util.ParseMoneyCN(td.Text()))
			})
			index++
		}
	})

	label := ""
	trlist := doc.Find("#BBMX_table tr")
	if trlist.Length() == 60 {
		trlist.Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				tr.Find("th.tips-fieldname-Right").Each(func(k int, th *goquery.Selection) {
					date := fmt.Sprint("20", th.Text())
					zcfzbList = append(zcfzbList, &Zcfzb{Date: date})
					lrbList = append(lrbList, &Lrb{Date: date})
					xjllbList = append(xjllbList, &Xjllb{Date: date})
				})
				label = "资产负债表"
				index = 0
				return
			}

			if tr.Find("th.tips-colname-Left").Length() > 0 {
				label = tr.Find("th").First().Text()
				index = 0
				return
			}

			tr.Find("td.tips-data-Right").Each(func(k int, td *goquery.Selection) {
				switch label {
				case "资产负债表":
					// fmt.Println(label, util.ParseMoneyCN(td.Text()), index, k)
					reflect.ValueOf(zcfzbList[k]).
						Elem().Field(index + 1).SetFloat(util.ParseMoneyCN(td.Text()))
				case "利润表":
					reflect.ValueOf(lrbList[k]).
						Elem().Field(index + 1).SetFloat(util.ParseMoneyCN(td.Text()))
				case "现金流量表":
					reflect.ValueOf(xjllbList[k]).
						Elem().Field(index + 1).SetFloat(util.ParseMoneyCN(td.Text()))
				}
			})
			index++

		})
	}

	trlist = doc.Find("#PPTable tr")
	if trlist.Length() == 20 {
		trlist.Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				tr.Find("th.tips-dataC").Each(func(k int, th *goquery.Selection) {
					finPercentList = append(finPercentList, &FinPercent{Date: fmt.Sprint("20", th.Text())})
				})
				return
			}

			if i == 1 {
				return
			}
			index = 0
			tr.Find("td.tips-data-Right").Each(func(k int, td *goquery.Selection) {
				if k%2 == 0 {
					return
				}
				reflect.ValueOf(finPercentList[index]).
					Elem().Field(i - 1).SetFloat(util.ParsePercent(td.Text()))
				index++
			})

		})
	}

	return cwzbList, zcfzbList, lrbList, xjllbList, finPercentList, nil
}

type Cwzb struct {
	Date string `json:"date"`
	// 每股指标
	A1 float64 `json:"a1"` // 基本每股收益(元)
	A2 float64 `json:"a2"` // 扣非每股收益(元)
	A3 float64 `json:"a3"` // 稀释每股收益(元)
	A4 float64 `json:"a4"` // 每股净资产(元)
	A5 float64 `json:"a5"` // 每股公积金(元)
	A6 float64 `json:"a6"` // 每股未分配利润(元)
	A7 float64 `json:"a7"` // 每股经营现金流(元)
	// 成长能力指标
	A8  float64 `json:"a8"`  // 营业收入(元)
	A9  float64 `json:"a9"`  // 毛利润(元)
	A10 float64 `json:"a10"` // 归属净利润(元)
	A11 float64 `json:"a11"` // 扣非净利润(元)
	A12 float64 `json:"a12"` // 营业收入同比增长(%)
	A13 float64 `json:"a13"` // 归属净利润同比增长(%)
	A14 float64 `json:"a14"` // 扣非净利润同比增长(%)
	A15 float64 `json:"a15"` // 营业收入滚动环比增长(%)
	A16 float64 `json:"a16"` // 归属净利润滚动环比增长(%)
	A17 float64 `json:"a17"` // 扣非净利润滚动环比增长(%)
	// 盈利能力指标
	A18 float64 `json:"a18"` // 加权净资产收益率(%)
	A19 float64 `json:"a19"` // 摊薄净资产收益率(%)
	A20 float64 `json:"a20"` // 摊薄总资产收益率(%)
	A21 float64 `json:"a21"` // 毛利率(%)
	A22 float64 `json:"a22"` // 净利率(%)
	A23 float64 `json:"a23"` // 实际税率(%)
	// 盈利质量指标
	A24 float64 `json:"a24"` // 预收款/营业收入
	A25 float64 `json:"a25"` // 销售现金流/营业收入
	A26 float64 `json:"a26"` // 经营现金流/营业收入
	// 运营能力指标
	A27 float64 `json:"a27"` // 总资产周转率(次)
	A28 float64 `json:"a28"` // 应收账款周转天数(天)
	A29 float64 `json:"a29"` // 存货周转天数(天)
	// 财务风险指标
	A30 float64 `json:"a30"` // 资产负债率(%)
	A31 float64 `json:"a31"` // 流动负债/总负债(%)
	A32 float64 `json:"a32"` // 流动比率
	A33 float64 `json:"a33"` // 速动比率

	StockID int64 `json:"stock_id" skip:"update"`
}

type Zcfzb struct {
	Date string  `json:"date"`
	A1   float64 `json:"a1"`  // 资产:货币资金(元)
	A2   float64 `json:"a2"`  // 应收账款(元)
	A3   float64 `json:"a3"`  // 其它应收款(元)
	A4   float64 `json:"a4"`  // 存货(元)
	A5   float64 `json:"a5"`  // 流动资产合计(元)
	A6   float64 `json:"a6"`  // 长期股权投资(元)
	A7   float64 `json:"a7"`  // 累计折旧(元)
	A8   float64 `json:"a8"`  // 固定资产(元)
	A9   float64 `json:"a9"`  // 无形资产(元)
	A10  float64 `json:"a10"` // 资产总计(元)
	A11  float64 `json:"a11"` // 负债:应付账款(元)
	A12  float64 `json:"a12"` // 预收账款(元)
	A13  float64 `json:"a13"` // 存货跌价准备(元)
	A14  float64 `json:"a14"` // 流动负债合计(元)
	A15  float64 `json:"a15"` // 长期负债合计(元)
	A16  float64 `json:"a16"` // 负债合计(元)
	A17  float64 `json:"a17"` // 权益:实收资本(或股本)(元)
	A18  float64 `json:"a18"` // 资本公积金(元)
	A19  float64 `json:"a19"` // 盈余公积金(元)
	A20  float64 `json:"a20"` // 股东权益合计(元)
	A21  float64 `json:"a21"` // 流动比率

	StockID int64 `json:"stock_id" skip:"update"`
}

type Lrb struct {
	Date string  `json:"date"`
	A1   float64 `json:"a1"`  // 营业收入(元)
	A2   float64 `json:"a2"`  // 营业成本(元)
	A3   float64 `json:"a3"`  // 销售费用(元)
	A4   float64 `json:"a4"`  // 财务费用(元)
	A5   float64 `json:"a5"`  // 管理费用(元)
	A6   float64 `json:"a6"`  // 资产减值损失(元)
	A7   float64 `json:"a7"`  // 投资收益(元)
	A8   float64 `json:"a8"`  // 营业利润(元)
	A9   float64 `json:"a9"`  // 利润总额(元)
	A10  float64 `json:"a10"` // 所得税(元)
	A11  float64 `json:"a11"` // 归属母公司所有者净利润(元)

	StockID int64 `json:"stock_id" skip:"update"`
}

type Xjllb struct {
	Date string  `json:"date"`
	A1   float64 `json:"a1"`  // 经营:销售商品、提供劳务收到的现金(元)
	A2   float64 `json:"a2"`  // 收到的税费返还(元)
	A3   float64 `json:"a3"`  // 收到其他与经营活动有关的现金(元)
	A4   float64 `json:"a4"`  // 经营活动现金流入小计(元)
	A5   float64 `json:"a5"`  // 购买商品、接受劳务支付的现金(元)
	A6   float64 `json:"a6"`  // 支付给职工以及为职工支付的现金(元)
	A7   float64 `json:"a7"`  // 支付的各项税费(元)
	A8   float64 `json:"a8"`  // 支付其他与经营活动有关的现金(元)
	A9   float64 `json:"a9"`  // 经营活动现金流出小计(元)
	A10  float64 `json:"a10"` // 经营活动产生的现金流量净额(元)
	A11  float64 `json:"a11"` // 投资:取得投资收益所收到的现金(元)
	A12  float64 `json:"a12"` // 处置固定资产、无形资产和其他长期...
	A13  float64 `json:"a13"` // 投资活动现金流入小计(元)
	A14  float64 `json:"a14"` // 购建固定资产、无形资产和其他长期...
	A15  float64 `json:"a15"` // 处置固定资产、无形资产和其他长期...
	A16  float64 `json:"a16"` // 投资支付的现金(元)
	A17  float64 `json:"a17"` // 投资活动现金流出小计(元)
	A18  float64 `json:"a18"` // 投资活动产生的现金流量净额(元)
	A19  float64 `json:"a19"` // 筹资:吸收投资收到的现金(元)
	A20  float64 `json:"a20"` // 取得借款收到的现金(元)
	A21  float64 `json:"a21"` // 筹资活动现金流入小计(元)
	A22  float64 `json:"a22"` // 偿还债务支付的现金(元)
	A23  float64 `json:"a23"` // 分配股利、利润或偿付利息支付的现...
	A24  float64 `json:"a24"` // 筹资活动现金流出小计(元)
	A25  float64 `json:"a25"` // 筹资活动产生的现金流量净额(元)

	StockID int64 `json:"stock_id" skip:"update"`
}

type FinPercent struct {
	Date string  `json:"date"`
	A1   float64 `json:"a1"`  // 营业收入(元)
	A2   float64 `json:"a2"`  // 营业成本(元)
	A3   float64 `json:"a3"`  // 营业税金及附加(元)
	A4   float64 `json:"a4"`  // 期间费用(元)
	A5   float64 `json:"a5"`  // 销售费用(元)
	A6   float64 `json:"a6"`  // 管理费用(元)
	A7   float64 `json:"a7"`  // 财务费用(元)
	A8   float64 `json:"a8"`  // 资产减值损失(元)
	A9   float64 `json:"a9"`  // 其他经营收益(元)
	A10  float64 `json:"a10"` // 公允价值变动损益(元)
	A11  float64 `json:"a11"` // 投资收益(元)
	A12  float64 `json:"a12"` // 营业利润(元)
	A13  float64 `json:"a13"` // 加:营业外收入(元)
	A14  float64 `json:"a14"` // 补贴收入(元)
	A15  float64 `json:"a15"` // 减:营业外支出(元)
	A16  float64 `json:"a16"` // 利润总额(元)
	A17  float64 `json:"a17"` // 减:所得税(元)
	A18  float64 `json:"a18"` // 净利润(元)

	StockID int64 `json:"stock_id" skip:"update"`
}
