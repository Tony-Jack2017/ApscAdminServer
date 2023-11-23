package router

import (
	"ApscAdmin/router/middleware"
	"ApscAdmin/router/routes"
	"ApscAdmin/router/routes/server"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api", middleware.CorsMiddleware())
	router.GET("/ws")
	RegisterRoutes(api)
	return router
}

func RegisterRoutes(group *gin.RouterGroup) {
	back := group.Group("/back")
	client := group.Group("/cl")
	common := group.Group("/common")
	RegisterBackRoutes(back)
	RegisterClientRoutes(client)
	RegisterCommonRoutes(common)
}

func RegisterCommonRoutes(common *gin.RouterGroup) {
	routes.CommonRoutes(common)
}

func RegisterClientRoutes(client *gin.RouterGroup) {
	routes.UserRoutes(client)
}

func RegisterBackRoutes(back *gin.RouterGroup) {

	admin := back.Group("/admin")
	serverGroup := back.Group("/server")

	// normal
	routes.AdminRoutes(admin)

	// server
	website := serverGroup.Group("/website")
	server.WebsiteRoutes(website)
}
