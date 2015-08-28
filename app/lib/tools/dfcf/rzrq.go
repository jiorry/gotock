package dfcf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/kere/gos"
	"github.com/kere/gos/lib/util"
)

// RzrqSumJSONData 融资融券汇总
type RzrqSumJSONData []string

// RzrqSumItemData 融资融券汇总
type RzrqSumItemData struct {
	Date     string `json:"date"`
	SHrzye   int64  `json:"sh_rzye"`
	SZrzye   int64  `json:"sz_rzye"`
	SMrzye   int64  `json:"sm_rzye"`
	SHrzmre  int64  `json:"sh_rzmre"`
	SZrzmre  int64  `json:"sz_rzmre"`
	SMrzmre  int64  `json:"sm_rzmre"`
	SHrqylye int64  `json:"sh_rqylye"`
	SZrqylye int64  `json:"sz_rqylye"`
	SMrqylye int64  `json:"sm_rqylye"`
	SHrzrqye int64  `json:"sh_rzrqye"`
	SZrzrqye int64  `json:"sz_rzrqye"`
	SMrzrqye int64  `json:"sm_rzrqye"`
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
	for _, item := range r {
		tmp = strings.Split(item, ",")
		if len(tmp) != 13 {
			return nil, fmt.Errorf("parse data error")
		}

		// t, err := time.Parse("2006-01-02", tmp[0])
		// if err != nil {
		// 	return nil, gos.DoError(err)
		// }
		itemData = &RzrqSumItemData{}
		itemData.Date = tmp[0]

		itemData.SHrzye, err = util.Str2Int64(tmp[1])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SZrzye, err = util.Str2Int64(tmp[2])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SMrzye, err = util.Str2Int64(tmp[3])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SHrzmre, err = util.Str2Int64(tmp[4])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SZrzmre, err = util.Str2Int64(tmp[5])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SMrzmre, err = util.Str2Int64(tmp[6])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SHrqylye, err = util.Str2Int64(tmp[7])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SZrqylye, err = util.Str2Int64(tmp[8])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SMrqylye, err = util.Str2Int64(tmp[9])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SHrzrqye, err = util.Str2Int64(tmp[10])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SZrzrqye, err = util.Str2Int64(tmp[11])
		if err != nil {
			return nil, gos.DoError(err)
		}
		itemData.SMrzrqye, err = util.Str2Int64(tmp[12])
		if err != nil {
			return nil, gos.DoError(err)
		}

		dataSet = append(dataSet, itemData)
	}

	return dataSet, nil
}

var sumdataCached []*RzrqSumItemData

// RzrqSumData 抓取数据
func RzrqSumData() ([]*RzrqSumItemData, error) {
	if sumdataCached != nil {
		return sumdataCached, nil
	}
	src, err := FetchRzrqSumData("SHSZHSSUM", 1)
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
func FetchRzrqSumData(sty string, page int) ([]byte, error) {
	st := time.Now().Unix() / 30
	formt := "http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=FD&sty=%s&st=0&sr=1&p=%d&ps=50&js=var%%20ruOtumOo={pages:(pc),data:[(x)]}&rt=%d"

	url := fmt.Sprintf(formt, sty, page, st)

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

	return body[28 : len(body)-1], nil
}
