package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"bufio"
	"net/http"
)

func main() {
	// s := "http://table.finance.yahoo.com/table.csv?s=601857.ss&c=2015&a=5&b=1&d=5&e=29&f=2015"
	s := "http://tmeet.cn"
	resp, err := http.Get(s)

	if err != nil {
		fmt.Println("error: ", err)
	} else if resp.Body == nil {
		fmt.Println("error: resp.Body is empty")
	} else {
		defer resp.Body.Close()
		// b, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(b))
		var line []byte
		var isPrefix bool
		var bf = bufio.NewReader(resp.Body.(io.Reader))

		for err == nil {
			line, isPrefix, err = bf.ReadLine()
			fmt.Println(string(line), isPrefix, err)
		}

		fmt.Println("err:---------------", err)
	}
}
