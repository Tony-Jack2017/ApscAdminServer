package routes

import (
	"ApscAdmin/app/controller"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(admin *gin.RouterGroup) {
	admin.POST("/create", controller.AdminCreate)
	admin.POST("/update", controller.AdminUpdate)
}
