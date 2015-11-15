package cninfo

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
)

func init() {

}

// FetchRzrqSumData 抓取数据
func FetchFinancialData(code, date string) ([]byte, error) {
	if len(date) != 10 {
		return nil, fmt.Errorf("date format error")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	formt := "http://www.cninfo.com.cn/information/stock/financialreport_.jsp?stockCode=%s&key=%s"

	u := fmt.Sprintf(formt, code, fmt.Sprint(r.Float64()))

	// resp, err := wget.Post(u, url.Values{"yyyy": {date[:4]}, "mm": {date[4:]}, "cwzb": {"financialreport"}, "button2": {"%CC%E1%BD%BB"}})
	// if err != nil {
	// 	return nil, err
	// }
	//
	// doc, err := goquery.NewDocumentFromResponse(resp)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// doc.Find("table td").Each(func(i int, s *goquery.Selection) {
	// 	if v, isok := s.Attr("bgcolor"); isok && v == "#daf2ff" {
	// 		fmt.Println(s.Text())
	// 	}
	// })

	body, err := wget.PostBody(u, url.Values{"yyyy": {date[:4]}, "mm": {date[4:]}, "cwzb": {"financialreport"}, "button2": {"%CC%E1%BD%BB"}})
	if err != nil {
		return nil, err
	}

	fmt.Println("--------------", date[:4], date[4:])
	fmt.Println(string(body))

	return nil, nil
}
