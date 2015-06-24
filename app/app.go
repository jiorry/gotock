package main

import (
	"./api"
	"./page/stock"

	"github.com/kere/gos"
	_ "github.com/lib/pq"
)

func main() {
	gos.Init()

	gos.Route("/", &stock.Main{})
	gos.RegRoute("^/stock/.+", &stock.Main{})

	// open api router
	gos.WebApiRoute("web", &api.Public{})

	// open api
	// api.RegistOpenApi()
	gos.WebApiRoute("open", &api.OpenApi{})

	// websocket router
	// gos.WebSocketRoute("conn", (*hiuser.UserWebSock)(nil))

	gos.Start()
}
