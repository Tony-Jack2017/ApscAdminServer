package main

import (
	"ApscAdmin/common/config"
	"ApscAdmin/common/server"
	"ApscAdmin/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

var HTTPSERVER *gin.Engine

func init() {
	config.ReadConfig()
	server.RunService()
	HTTPSERVER = router.SetupRouter()
}

func main() {
	HTTPSERVER.Run(fmt.Sprintf("%s:%d", config.Conf.System.Host, config.Conf.System.HttpPort))
}
