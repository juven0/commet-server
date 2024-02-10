package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
}

func (Server) StartServer(routes *gin.Engine) {
	// port := os.Getenv("PORT")
	// host := os.Getenv("HOST")
	// if port != "" && host != "" {
	// 	routes.Run(host + ":" + port)
	// } else {
	// 	routes.Run()
	// }
	routes.Run()
}
