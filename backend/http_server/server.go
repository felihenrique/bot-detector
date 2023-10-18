package http_server

import (
	"botdetector/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	server := gin.Default()

	server.GET("/requests", Controllers.Read)
	server.POST("/requests", Controllers.Write)

	err := server.Run("0.0.0.0:" + config.Config.Port)

	if err != nil {
		panic("Can't start the http server: " + err.Error())
	}
}
