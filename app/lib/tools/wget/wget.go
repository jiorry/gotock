package wget

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kere/gos"
)

// Wget 抓取数据
func Get(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36`)
	resp, err := client.Do(req)

	if err != nil {
		return nil, gos.DoError("error:", err)
	} else if resp.Body == nil {
		return nil, gos.DoError("error: resp.Body is empty")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, gos.DoError(err)
	}

	if resp.StatusCode != 200 {
		return nil, gos.DoError(fmt.Sprintf("wget failed:%d", resp.StatusCode))
	}
	// var OKPJKmpr={pages:0,data:[{stats:false}]}
	return body, nil
}
