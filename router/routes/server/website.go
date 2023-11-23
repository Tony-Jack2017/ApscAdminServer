package server

import (
	"ApscAdmin/app/controller/server"
	"github.com/gin-gonic/gin"
)

func WebsiteRoutes(website *gin.RouterGroup) {
	website.POST("/add", server.WebsiteAdd)
}
