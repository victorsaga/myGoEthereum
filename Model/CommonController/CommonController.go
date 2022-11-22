package CommonController

import (
	"encoding/json"
	"fmt"
	"io"
	"myGoEthereum/Helper/TimeHelper"
	"myGoEthereum/Model/CommonModel"
	"myGoEthereum/Model/ResultCode"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CanBindJson(c *gin.Context, i interface{}) bool {
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(200, CommonModel.ApiResponse{
			Success: false,
			Code:    "90000",
			Message: fmt.Sprint(err),
		})
		return false
	}
	return true
}

func SetDefaultValue(i interface{}) interface{} {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	t := CommonModel.ApiResponseWithData{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		panic(err)
	}

	if t.Code == "" {
		t.Success = true
		t.Code = ResultCode.Success
	}

	return t
}

func ServeExcelContent(fileName string, excelContent io.ReadSeeker, c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(c.Writer, c.Request, fileName, TimeHelper.GetUTC8TimeNow(), excelContent)
}
