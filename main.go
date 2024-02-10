package main

import (
	"comete-server/internal/routes"
	"comete-server/internal/server"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier .env:", err)
		return
	}
	routes := routes.Routes()
	var cometeServer server.Server
	cometeServer.StartServer(routes)

}
