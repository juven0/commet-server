package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
}

func (Server) StartServer(routes *gin.Engine) {
	port := os.Getenv("PORT")
	routes.Run("localhost:" + port)
}
