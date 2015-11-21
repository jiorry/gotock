package dfcf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos"
	"github.com/kere/gos/db"
)

// FillYyb 抓取数据
func FillYyb() error {
	pages, err := fetchAndFillYyb(1)
	if err != nil {
		return err
	}
	for i := 2; i < pages+1; i++ {
		_, err = fetchAndFillYyb(i)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// fetchAndFillYyb 抓取数据
func fetchAndFillYyb(page int) (int, error) {
	// http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=LHB&sty=YYTJ&stat=6&sr=0&st=1&p=2&ps=50&js=var%20XvAVhGPE={%22data%22:[(x)],%22pages%22:%22(pc)%22,%22update%22:%22(ud)%22}&rt=48257541
	formt := `http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=LHB&sty=YYTJ&stat=6&sr=0&st=1&p=%d&ps=%d&js=var%%20XvAVhGPE={"data":[(x)],"pages":"(pc)","update":"(ud)"}&rt=%d`
	pageLimit := 100

	body, err := wget.GetBody(fmt.Sprintf(formt, page, pageLimit, time.Now().Unix()))
	if err != nil {
		return -1, err
	}

	exists := db.NewExistsBuilder("yyb")
	ist := db.NewInsertBuilder("yyb")
	upd := db.NewUpdateBuilder("yyb")
	query := db.NewQueryBuilder("yyb")
	var yyb *yybStruct
	var yybLub *yybLhbStruct
	var arr []string
	var row db.DataRow
	src := body[bytes.Index(body, []byte("={"))+1:]

	v := &jsonYybData{}
	err = json.Unmarshal(src, v)
	if err != nil {
		return -1, err
	}
	// "80136686,2,216,135,七喜控股|协鑫集成|恒信移动,2991067145.735,广东,80000073,440000,105,6089657316.805,3098590171.07,华泰证券股份有限公司广州天河东路证券营业部,002027.SZ|002506.SZ|300081.SZ"
	for i, item := range v.Data {
		arr = strings.Split(item, ",")
		yyb = &yybStruct{Code: arr[0], Name: arr[12], Area: arr[6]}

		if !exists.Table("yyb").Where("code=?", yyb.Code).Exists() {
			gos.Log.Info("insert", yyb.Code)
			ist.Table("yyb").Insert(yyb)
		}

		row, _ = query.Table("yyb").Where("code=?", yyb.Code).QueryOne()
		if row.Empty() {
			return -1, fmt.Errorf("code %s not found", yyb.Code)
		}

		yybLub = &yybLhbStruct{
			YybID:   row.GetInt64("id"),
			Updated: v.Update,
			Rank:    pageLimit*(page-1) + i + 1,
			Amount:  util.ParseMoney(arr[10]),
			Buy:     util.ParseMoney(arr[11]),
			Sell:    util.ParseMoney(arr[5]),
			Num:     int(util.ParseMoney(arr[2])),
			NumBuy:  int(util.ParseMoney(arr[3])),
			NumSell: int(util.ParseMoney(arr[9])),
		}

		if exists.Table("yyb_lhb").Where("yyb_id=?", yybLub.YybID).Exists() {
			gos.Log.Info("update", yyb.Code, yybLub.YybID)
			upd.Table("yyb_lhb").Where("yyb_id=?", yybLub.YybID).Update(yybLub)
		} else {
			gos.Log.Info("insert", yyb.Code, yybLub.YybID)
			ist.Table("yyb_lhb").Insert(yybLub)
		}

	}
	return int(util.ParseMoney(v.Pages)), nil
}

// 营业部
type yybStruct struct {
	ID   int64  `json:"id" skip:"all"`
	Name string `json:"name"`
	Code string `json:"code"`
	Area string `json:"area"` // 所在地区
}

type yybLhbStruct struct {
	YybID   int64   `json:"yyb_id" skip:"update"`
	Updated string  `json:"updated"`
	Rank    int     `json:"rank"`     // 排名
	Amount  float64 `json:"amount"`   // 龙虎榜成交金额
	Buy     float64 `json:"buy"`      //
	Sell    float64 `json:"sell"`     //
	Num     int     `json:"num"`      // 龙虎榜上榜次数
	NumBuy  int     `json:"num_buy"`  // 买入上榜次数
	NumSell int     `json:"num_sell"` // 卖出上榜次数

}
type jsonYybData struct {
	Pages  string   `json:"pages"`
	Update string   `json:"update"`
	Data   []string `json:"data"`
}
