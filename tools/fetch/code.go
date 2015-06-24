package main

import (
	"fmt"
	"github.com/crufter/goquery"
)

func main() {
	s := "http://quote.eastmoney.com/stocklist.html"
	x, err := goquery.ParseUrl(s)
	if err != nil {
		panic(err)
	}

	nodes := x.Find("#quotesearch ul li")
	arr := nodes.HtmlAll()
	for _, item := range arr {
		fmt.Println(item)
	}
}
