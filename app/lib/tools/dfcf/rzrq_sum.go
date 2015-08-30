package dfcf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/kere/gos"
	"github.com/kere/gos/lib/log"
	"github.com/kere/gos/lib/util"
)

var rzrqSumDataMapping []*sumDataMapping

type sumDataMapping struct {
	Name  string
	Index int
}

func init() {
	rzrqSumDataMapping = make([]*sumDataMapping, 12)
	rzrqSumDataMapping[0] = &sumDataMapping{"SHrzye", 1}
	rzrqSumDataMapping[1] = &sumDataMapping{"SZrzye", 2}
	rzrqSumDataMapping[2] = &sumDataMapping{"SMrzye", 3}

	rzrqSumDataMapping[3] = &sumDataMapping{"SHrzmre", 4}
	rzrqSumDataMapping[4] = &sumDataMapping{"SZrzmre", 5}
	rzrqSumDataMapping[5] = &sumDataMapping{"SMrzmre", 6}

	rzrqSumDataMapping[6] = &sumDataMapping{"SHrqylye", 7}
	rzrqSumDataMapping[7] = &sumDataMapping{"SZrqylye", 8}
	rzrqSumDataMapping[8] = &sumDataMapping{"SMrqylye", 9}

	rzrqSumDataMapping[9] = &sumDataMapping{"SHrzrqye", 10}
	rzrqSumDataMapping[10] = &sumDataMapping{"SZrzrqye", 11}
	rzrqSumDataMapping[11] = &sumDataMapping{"SMrzrqye", 12}
}

// RzrqSumJSONData 融资融券汇总
type RzrqSumJSONData []string

// RzrqSumItemData 融资融券汇总
type RzrqSumItemData struct {
	Date   time.Time `json:"date"`
	SHrzye int64     `json:"sh_rzye"`
	SZrzye int64     `json:"sz_rzye"`
	SMrzye int64     `json:"sm_rzye"`

	SHrzmre int64 `json:"sh_rzmre"`
	SZrzmre int64 `json:"sz_rzmre"`
	SMrzmre int64 `json:"sm_rzmre"`

	SHrqylye int64 `json:"sh_rqylye"`
	SZrqylye int64 `json:"sz_rqylye"`
	SMrqylye int64 `json:"sm_rqylye"`

	SHrzrqye int64 `json:"sh_rzrqye"`
	SZrzrqye int64 `json:"sz_rzrqye"`
	SMrzrqye int64 `json:"sm_rzrqye"`
}

// ParseSumData 解析两市汇总信息
func (r RzrqSumJSONData) ParseSumData() ([]*RzrqSumItemData, error) {
	if len(r) == 0 {
		return nil, nil
	}
	var err error
	var dataSet = make([]*RzrqSumItemData, 0)

	var tmp []string
	var itemData *RzrqSumItemData
	var x int64
	var val reflect.Value

	for _, item := range r {
		tmp = strings.Split(item, ",")
		if len(tmp) != 13 {
			return nil, fmt.Errorf("parse data error")
		}

		itemData = &RzrqSumItemData{}
		itemData.Date, err = time.Parse("2006-01-02", tmp[0])
		if err != nil {
			return nil, gos.DoError(err)
		}

		val = reflect.ValueOf(itemData).Elem()
		for _, mapping := range rzrqSumDataMapping {
			if tmp[mapping.Index] == "-" || tmp[mapping.Index] == "" {
				val.FieldByName(mapping.Name).SetInt(-1)
				continue
			}

			if x, err = util.Str2Int64(tmp[mapping.Index]); err != nil {
				return nil, gos.DoError(err)
			}
			val.FieldByName(mapping.Name).SetInt(x)
		}
		dataSet = append(dataSet, itemData)
	}

	// remove -1 value
	for i, item := range dataSet {
		val = reflect.ValueOf(item).Elem()
		for _, mapping := range rzrqSumDataMapping {
			if val.FieldByName(mapping.Name).Int() == -1 {
				var tmpIndex = i - 1
				if i == 0 {
					tmpIndex = i + 1
				}
				val.FieldByName(mapping.Name).SetInt(reflect.ValueOf(dataSet[tmpIndex]).Elem().FieldByName(mapping.Name).Int())
			}
		}
	}

	return dataSet, nil
}

var sumdataCached []*RzrqSumItemData

func isCached() bool {
	if len(sumdataCached) == 0 {
		return false
	}

	t := sumdataCached[0].Date
	df := "20060102"
	now := time.Now()
	tStr := t.Format(df)

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
		if t.Format(df) == tStr {
			return true
		}
	}
	return false
}

// RzrqSumData 抓取数据
func GetRzrqSumData() ([]*RzrqSumItemData, error) {
	if isCached() {
		log.App.Info("rzrq sum cached")
		return sumdataCached, nil
	}

	src, err := FetchRzrqSumData(1)
	if err != nil {
		return nil, gos.DoError(err)
	}
	v := &RzrqSumJSONData{}
	if err = json.Unmarshal(src, &v); err != nil {
		return nil, gos.DoError(err)
	}

	sumdataCached, err = v.ParseSumData()
	return sumdataCached, err
}

// FetchRzrqSumData 抓取数据
func FetchRzrqSumData(page int) ([]byte, error) {
	st := time.Now().Unix() / 30
	formt := "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=%s&st=0&sr=1&p=%d&ps=50&js=var%%20ruOtumOo={pages:(pc),data:[(x)]}&rt=%d"

	url := fmt.Sprintf(formt, "SHSZHSSUM", page, st)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36`)
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	} else if resp.Body == nil {
		return nil, gos.DoError("error: resp.Body is empty")
	}

	log.App.Info("rzrq fetch sum data")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, gos.DoError(err)
	}

	if resp.StatusCode != 200 {
		return nil, gos.DoError(fmt.Sprintf("http get failed:%d", resp.StatusCode))
	}

	return body[bytes.Index(body, []byte("[")) : len(body)-1], nil
}
