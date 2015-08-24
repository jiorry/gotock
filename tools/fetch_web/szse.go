package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("ACTIONID", "7")
	v.Add("AJAX", "AJAX-TRUE")
	v.Add("CATALOGID", "1803")
	v.Add("TABKEY", "tab1")
	v.Add("txtQueryDate", "2015-05-29")
	v.Add("REPORT_ACTION", "search")
	strURL := "http://www.szse.cn/main/marketdata/tjsj/jbzb/"

	resp, err := http.PostForm(strURL, v)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
