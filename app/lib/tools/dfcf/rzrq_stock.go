package dfcf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/kere/gos"
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
	for _, item := range r {
		tmp = strings.Split(item, ",")

		itemData = &RzrqStockData{}
		itemData.Date, err = time.Parse("2006/1/2", tmp[4])
		if err != nil {
			return nil, gos.DoError(err)
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
	src, err := FetchRzrqStockData(code, 1)
	if err != nil {
		return nil, gos.DoError(err)
	}

	v := &RzrqStockJSONData{}
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
	//var% OKPJKmpr={pages:10,data:
	//http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=MTE&mkt=1&code=600718&st=0&sr=1&p=5&ps=50&js=var%20OKPJKmpr={pages:(pc),data:[(x)]}&rt=48027423
	formt := "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=MTE&mkt=1&code=%s&st=0&sr=1&p=%d&ps=50&js=var%%20OKPJKmpr={pages:(pc),data:[(x)]}&rt=%d"

	url := fmt.Sprintf(formt, code, page, st)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36`)
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	} else if resp.Body == nil {
		return nil, gos.DoError("error: resp.Body is empty")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, gos.DoError(err)
	}

	if resp.StatusCode != 200 {
		return nil, gos.DoError(fmt.Sprintf("http get failed:%d", resp.StatusCode))
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
	now := time.Now()
	nowStr := now.Format(df)

	switch now.Weekday() {
	case time.Sunday:
		if now.AddDate(0, 0, -1).Format(df) == nowStr {
			return true
		}
	case time.Monday:
		if now.AddDate(0, 0, -2).Format(df) == nowStr {
			return true
		}
	default:
		if t.Format(df) == nowStr {
			return true
		}
	}
	return false
}
