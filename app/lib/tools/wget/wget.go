package wget

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kere/gos"
)

// Wget 抓取数据
func Get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36`)
	resp, err := client.Do(req)

	if err != nil {
		return nil, gos.DoError("error:", err)
	} else if resp.Body == nil {
		return nil, gos.DoError("error: resp.Body is empty")
	}

	if resp.StatusCode != 200 {
		return resp, gos.DoError(fmt.Sprintf("Get failed:%d", resp.StatusCode))
	}
	// var OKPJKmpr={pages:0,data:[{stats:false}]}
	return resp, nil
}

func GetBody(url string) ([]byte, error) {
	resp, err := Get(url)
	if err != nil {
		return nil, gos.DoError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, gos.DoError(err)
	}

	return body, nil
}

func Post(url string, val url.Values) (*http.Response, error) {
	resp, err := http.PostForm(url, val)
	if err != nil {
		return nil, gos.DoError(err)
	}

	if resp.StatusCode != 200 {
		return resp, gos.DoError(fmt.Sprintf("Post failed:%d", resp.StatusCode))
	}

	return resp, nil
}

func PostBody(url string, val url.Values) ([]byte, error) {
	resp, err := Post(url, val)
	if err != nil {
		return nil, gos.DoError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, gos.DoError(err)
	}

	return body, nil
}
