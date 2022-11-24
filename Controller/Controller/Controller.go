package Controller

import (
	"myGoEthereum/Model/BaseModel"
	"myGoEthereum/Model/CommonController"
	"myGoEthereum/Service/Service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// root
func GetBlocksByLimit(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.GetBlocksByLimit(limit)))
}

func GetBlockTransactionHashes(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.GetBlockTransactionHashes(id)))
}

func GetTransactionsReceiptLogs(c *gin.Context) {
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.GetTransactionsReceiptLogs(c.Param("txHash"))))
}

// /api
func GetNewBlocks(c *gin.Context) {
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.GetNewBlocks()))
}

func InitialDbData(c *gin.Context) {
	c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.InitialDbData()))
}

func Login(c *gin.Context) {
	json := BaseModel.LoginRequest{}
	if CommonController.CanBindJson(c, &json) {
		c.JSON(http.StatusOK, CommonController.SetDefaultValue(Service.Login(json)))
	}
}
