package routes

import (
	"ApscAdmin/app/controller"
	"github.com/gin-gonic/gin"
)

func AccountRoutes(account *gin.RouterGroup) {
	account.POST("/update", controller.AccountUpdate)
}
