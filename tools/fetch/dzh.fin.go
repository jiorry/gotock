package main

import (
	"../../app/lib/tools/dzh"
	"fmt"
)

func main() {
	name := `C:\Program360\dzh2_hlzq\Download\FIN\full.FIN`
	_, err := dzh.ReadData(name)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
}
