package main

import (
	"comete-server/internal/routes"
	"comete-server/internal/server"
)

func main() {
	routes := routes.Routes()
	var cometeServer server.Server
	cometeServer.StartServer(routes)

}
