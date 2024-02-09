package server

import (
	"comete-server/configs"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
}

func (Server) StartServer(routes *gin.Engine) {
	var conf configs.Configs
	configuration := conf.LoadConfigs()
	routes.Run(configuration.Server.Host + ":" + configuration.Server.Port)
}
