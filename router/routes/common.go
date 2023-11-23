package routes

import (
	"ApscAdmin/app/controller"
	"github.com/gin-gonic/gin"
)

func CommonRoutes(common *gin.RouterGroup) {
	common.GET("/verify_code", controller.GetVerifyCode)
	common.POST("/single/upload_file", controller.UploadFileSingle)
	common.POST("/multiple/upload_file", controller.UploadFileMultiple)
}
