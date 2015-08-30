package main

import (
	"./api"
	"./page/home"
	"./page/rzrq"
	"./page/user"

	"github.com/kere/gos"
	_ "github.com/lib/pq"

	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	gos.Init()

	gos.Route("/", &home.Default{})
	gos.Route("/login", &user.Login{})
	gos.Route("/regist", &user.Regist{})

	gos.Route("/rzrq/sum", &rzrq.Sum{})
	gos.Route("/rzrq/stock/:code", &rzrq.Stock{})

	// open api router
	// gos.WebApiRoute("web", &api.Public{})

	// open api
	// api.RegistOpenApi()
	gos.WebApiRoute("open", &api.OpenApi{})

	// websocket router
	// gos.WebSocketRoute("conn", (*hiuser.UserWebSock)(nil))

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	gos.Start()
}
