package dfcf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
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

	return dataSet, nil
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
	st := gos.NowInLocation().Unix() / 30
	formt := "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=%s&st=0&sr=1&p=%d&ps=50&js=var%%20ruOtumOo={pages:(pc),data:[(x)]}&rt=%d"

	url := fmt.Sprintf(formt, "SHSZHSSUM", page, st)

	body, err := wget.GetBody(url)
	if err != nil {
		return nil, gos.DoError(err)
	}

	return body[bytes.Index(body, []byte("[")) : len(body)-1], nil
}

var sumdataCached []*RzrqSumItemData

func isCached() bool {
	if len(sumdataCached) == 0 {
		return false
	}

	t := sumdataCached[0].Date
	df := "20060102"
	now := gos.NowInLocation()

	// 00:00 - 08:00 当作前一天看待
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
			if sumdataCached[0].SMrzye < 0 {
				return false
			}

			return true
		}
	default:
		if nowStr == tStr {
			return true
		}
	}
	return false
}
