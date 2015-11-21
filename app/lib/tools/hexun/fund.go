package hexun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos/db"

	iconv "github.com/djimenez/iconv-go"
)

func FundType(ctype string) int {
	switch ctype {
	case "基金":
		return 1
	case "保险":
		return 2
	case "社保":
		return 3
	case "QFII":
		return 4
	case "券商":
		return 5
	case "信托":
		return 6
	case "银行":
		return 7
	case "投资公司":
		return 8
	default:
		return -1
	}
}

// FetchRzrqSumData 抓取数据
func FillFundRankAll(quarter string) error {
	pages, err := fetchAndFillFundRank(1, quarter)
	if err != nil {
		return err
	}
	for i := 2; i < pages+1; i++ {
		_, err = fetchAndFillFundRank(i, quarter)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func fetchAndFillFundRank(page int, quarter string) (int, error) {
	// http://stockdata.stock.hexun.com/jgcc/data/outdata/orgrank.ashx?count=50&date=2015-09-30&orgType=&stateType=null&titType=null&page=2&callback=hxbase_json7
	formt := "http://stockdata.stock.hexun.com/jgcc/data/outdata/orgrank.ashx?count=%s&date=%s&orgType=&stateType=null&titType=null&page=%d&callback=hxbase_json7"
	pageLimit := 100
	body, err := wget.GetBody(fmt.Sprintf(formt, pageLimit, quarter, page))
	if err != nil {
		return -1, err
	}

	exists := db.NewExistsBuilder("funds")
	ist := db.NewInsertBuilder("funds")
	query := db.NewQueryBuilder("funds")

	var fund *Fund
	var fundRank *FundRank
	var row db.DataRow

	// hxbase_json7(
	str, err := iconv.ConvertString(string(body), "gb2312", "utf-8")
	if err != nil {
		return -1, err
	}

	src := []byte(str)
	src = src[13 : len(src)-1]
	src = bytes.Replace(src, []byte(":'"), []byte(`":"`), -1)
	src = bytes.Replace(src, []byte("',"), []byte(`","`), -1)
	src = bytes.Replace(src, []byte("'}"), []byte(`"}`), -1)
	src = bytes.Replace(src, []byte("{"), []byte(`{"`), -1)
	src = bytes.Replace(src, []byte("sum:"), []byte(`sum":`), 1)
	src = bytes.Replace(src, []byte("list:"), []byte(`"list":`), 1)

	v := &JsonFund{}
	err = json.Unmarshal(src, v)
	if err != nil {
		fmt.Println(string(src))
		return -1, err
	}
	// {RankTd:'51',OrgName:'法国巴黎银行',OrgNameLink:'o-QF000031.shtml',OrgType:'QFII',ShareHoldingNum:'3',ShareHoldingNumLink:'otherDetail.aspx?OrgNo=QF000031',TotalHoldings:'48,388.00',TotalMarketValue:'1,100,735.00',OrgAlt:'法国巴黎银行'}
	for _, item := range v.List {
		fund = &Fund{Code: item["OrgNameLink"], Name: string(item["OrgName"]), TypeID: FundType(item["OrgType"])}
		fund.Code = fund.Code[2 : len(fund.Code)-6]
		fmt.Println("----")
		if !exists.Table("funds").Where("code=? and type_id=?", fund.Code, fund.TypeID).Exists() {
			fmt.Println("insert", fund)
			ist.Table("funds").Insert(fund)
		}

		row, _ = query.Table("funds").Where("code=? and type_id=?", fund.Code, fund.TypeID).QueryOne()
		if row.Empty() {
			return -1, fmt.Errorf("code %s not found", fund.Code)
		}

		fundRank = &FundRank{
			FundID: row.GetInt64("id"),
			Date:   quarter,
			Rank:   int(util.ParseMoney(item["RankTd"])),
			Count:  int(util.ParseMoney(item["ShareHoldingNum"])),
			MH:     int64(util.ParseMoney(item["TotalHoldings"])),
			MV:     int64(util.ParseMoney(item["TotalMarketValue"])),
		}

		if !exists.Table("fund_rank").Where("fund_id=? and date=?", fundRank.FundID, fundRank.Date).Exists() {
			fmt.Println("insert", fundRank)
			ist.Table("fund_rank").Insert(fundRank)
		}

	}
	return int(math.Ceil(float64(v.Sum / pageLimit))), nil
}

type Fund struct {
	ID     int64  `json:"id" skip:"all"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	TypeID int    `json:"type_id"` //
}

type FundRank struct {
	FundID int64  `json:"fund_id" skip:"update"`
	Date   string `json:"date"`
	Rank   int    `json:"rank"`  // 排名
	Count  int    `json:"count"` // 股票数量
	MV     int64  `json:"mv"`    // 持股总数（万股）
	MH     int64  `json:"mh"`    // 持股市值（万元）
}
type JsonFund struct {
	Sum  int                 `json:"sum"`
	List []map[string]string `json:"list"`
}
