package Controller

import (
	"myGoEthereum/Model/BaseModel"
	"myGoEthereum/Model/CommonController"
	"myGoEthereum/Service/Service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	json := BaseModel.LoginRequest{}
	if CommonController.CanBindJson(c, &json) {
		c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.Login(json)))
	}
}

func GetNewBlocks(c *gin.Context) {
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.GetNewBlocks()))
}

func InitialDbData(c *gin.Context) {
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.InitialDbData()))
}
