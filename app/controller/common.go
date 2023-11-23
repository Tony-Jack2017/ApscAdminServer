package controller

import (
	"ApscAdmin/app/service"
	"ApscAdmin/common/model"
	"github.com/gin-gonic/gin"
)

func GetVerifyCode(context *gin.Context) {
}

func UploadFileSingle(context *gin.Context) {
	file, _ := context.FormFile("file")
	err, msg := service.UploadFileSingleSVC(file)
	if err != nil {
		context.JSON(500, model.Response{
			Code:    "-1",
			Status:  "ERROR",
			Message: msg,
		})
		return
	} else {
		context.JSON(200, model.Response{
			Code:    "0",
			Status:  "SUCCESS",
			Message: msg,
		})
		return
	}
}

func UploadFileMultiple(context *gin.Context) {
}
