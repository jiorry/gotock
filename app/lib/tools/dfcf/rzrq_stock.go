package dfcf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/kere/gos"
	"github.com/kere/gos/lib/log"
	"github.com/kere/gos/lib/util"
)

var rzrqStockDataMapping []*stockDataMapping

type stockDataMapping struct {
	Name  string
	Index int
}

func init() {
	rzrqStockDataMapping = make([]*stockDataMapping, 9)
	rzrqStockDataMapping[0] = &stockDataMapping{"Rzye", 12}
	rzrqStockDataMapping[1] = &stockDataMapping{"Rqye", 7}
	rzrqStockDataMapping[2] = &stockDataMapping{"Rzmre", 10}
	rzrqStockDataMapping[3] = &stockDataMapping{"Rzche", 9}
	rzrqStockDataMapping[4] = &stockDataMapping{"Rzjme", 13}
	rzrqStockDataMapping[5] = &stockDataMapping{"Rqyl", 8}
	rzrqStockDataMapping[6] = &stockDataMapping{"Rqmcl", 6}
	rzrqStockDataMapping[7] = &stockDataMapping{"Rqchl", 5}
	rzrqStockDataMapping[8] = &stockDataMapping{"Rzrqye", 11}

	stockDataCached = make(map[string][]*RzrqStockData, 0)
}

// RzrqStockData 融资融券
type RzrqStockJSONData []string

// RzrqStockData 融资融券汇总
// 600718,融资融券_沪证,东软集团,1265057757,2014/11/7,1258800.00,1257200,988640.1,68418,181606331.00,162010393,1267035037.1,1266046397,-19595938.00
type RzrqStockData struct {
	Code   string    `json:"code"`
	Type   string    `json:"type"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`   //4
	Rzye   float64   `json:"rzye"`   //12 融资余额
	Rqye   float64   `json:"rqye"`   //7 融券余额
	Rzmre  float64   `json:"rzmre"`  //10 融资买入额
	Rzche  float64   `json:"rzche"`  //9 融资偿还额
	Rzjme  float64   `json:"rzjme"`  //13 融资净买额
	Rqyl   float64   `json:"rqyl"`   //8 融券余量
	Rqmcl  float64   `json:"rqmcl"`  //6 融券卖出量
	Rqchl  float64   `json:"rqchl"`  //5 融券偿还量
	Rzrqye float64   `json:"rzrqye"` //11 融资融券余额
}

// ParseSumData 解析两市汇总信息
func (r RzrqStockJSONData) ParseSumData() ([]*RzrqStockData, error) {
	if len(r) == 0 {
		return nil, nil
	}
	var err error
	var dataSet = make([]*RzrqStockData, 0)

	var tmp []string
	var itemData *RzrqStockData
	var x float64
	var val reflect.Value
	var mapping *stockDataMapping
	var isFirst = true

	for _, item := range r {
		tmp = strings.Split(item, ",")

		itemData = &RzrqStockData{}
		itemData.Date, err = time.Parse("2006/1/2", tmp[4])
		if err != nil {
			return nil, gos.DoError(err)
		}

		if isFirst {
			itemData.Code = tmp[0]
			itemData.Name = tmp[2]
			isFirst = false
		}

		val = reflect.ValueOf(itemData).Elem()
		for _, mapping = range rzrqStockDataMapping {
			if tmp[mapping.Index] == "" || tmp[mapping.Index] == "-" {
				continue
			}

			if x, err = strconv.ParseFloat(tmp[mapping.Index], 64); err != nil {
				return nil, gos.DoError(err)
			}

			val.FieldByName(mapping.Name).SetFloat(x)
		}

		dataSet = append(dataSet, itemData)
	}

	return dataSet, nil
}

// GetRzrqStockData 抓取数据
func GetRzrqStockData(code string) ([]*RzrqStockData, error) {
	if isStockCached(code) {
		log.App.Info("rzrq stock cached", code)
		return stockDataCached[code], nil
	}

	src, err := FetchRzrqStockData(code, 1)
	if err != nil {
		return nil, gos.DoError(err)
	}

	v := &RzrqStockJSONData{}
	// [{stats:false}]
	if len(src) == 15 && string(src) == "[{stats:false}]" {
		return nil, fmt.Errorf("您所查找的股票代码 %s 不存在", code)
	}

	if err = json.Unmarshal(src, &v); err != nil {
		return nil, gos.DoError(err)
	}
	var dataSet []*RzrqStockData
	dataSet, err = v.ParseSumData()
	if err != nil {
		return nil, err
	}

	stockDataCached[code] = dataSet
	return dataSet, err
}

// FetchRzrqStockData 抓取数据
func FetchRzrqStockData(code string, page int) ([]byte, error) {
	st := time.Now().Unix() / 30
	// var% OKPJKmpr={pages:10,data:
	// http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=MTE&mkt=1&code=600718&st=0&sr=1&p=5&ps=50&js=var%20OKPJKmpr={pages:(pc),data:[(x)]}&rt=48027423
	// http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=MTE&mkt=2&code=002161&st=0&sr=1&p=5&ps=50&js=var%20QOwpoBrj={pages:(pc),data:[(x)]}&rt=48154207
	// http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=MTE&mkt=2&code=300001&st=0&sr=1&p=5&ps=50&js=var%20laheCxDp={pages:(pc),data:[(x)]}&rt=48154224

	formt := "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=MTE&mkt=%d&code=%s&st=0&sr=1&p=%d&ps=50&js=var%%20OKPJKmpr={pages:(pc),data:[(x)]}&rt=%d"
	mkt := 2
	switch code[0:1] {
	case "6":
		mkt = 1
	}

	url := fmt.Sprintf(formt, mkt, code, page, st)

	body, err := wget.GetBody(url)
	if err != nil {
		return nil, gos.DoError(err)
	}

	// var OKPJKmpr={pages:0,data:[{stats:false}]}
	return body[bytes.Index(body, []byte("[")) : len(body)-1], nil
}

var stockDataCached map[string][]*RzrqStockData

func StockCachedList() []string {
	l := make([]string, 0)
	for k, _ := range stockDataCached {
		l = append(l, k)
	}
	return l
}

func isStockCached(code string) bool {
	var isOk bool
	var v []*RzrqStockData

	if v, isOk = stockDataCached[code]; !isOk {
		return false
	}
	if len(v) == 0 {
		return false
	}

	t := v[0].Date
	df := "20060102"
	now := gos.NowInLocation()

	// 00:00 - 08:00
	if now.Hour() < 8 {
		now = now.AddDate(0, 0, -1)
	}

	tStr := t.Format(df)
	nowStr := now.Format(df)

	appConf := gos.Configuration.GetConf("other")
	if util.InStringSlice(appConf.GetStringSlice("holiday"), nowStr) {
		return true
	}

	switch now.Weekday() {
	case time.Sunday:
		if now.AddDate(0, 0, -2).Format(df) == tStr {
			return true
		}
	case time.Monday:
		if now.AddDate(0, 0, -3).Format(df) == tStr {
			return true
		}
	default:
		if now.Format(df) == tStr {
			return true
		}
	}
	return false
}
