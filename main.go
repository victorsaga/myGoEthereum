package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"myGoEthereum/Controller/Controller"
	"myGoEthereum/Helper/ConfigHelper"
	"myGoEthereum/Helper/LogHelper"
	"myGoEthereum/Helper/StructHelper"
	"myGoEthereum/Helper/TimeHelper"
	"myGoEthereum/Model/CommonModel"
	"myGoEthereum/Model/ResultCode"
	"myGoEthereum/Service/Service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	g := gin.New()
	g.Use(CustomLogger)
	g.Use(PanicHandler)
	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	//不用verifyJwtToken的路由
	api := g.Group("/api")
	{
		api.POST("/login", Controller.Login)
		api.POST("/initialDbData", Controller.InitialDbData)
		api.POST("/getNewBlocks", Controller.GetNewBlocks)
	}

	//用verifyJwtToken的路由
	// verifyToken := g.Group("/t").Use(verifyJwtToken)
	// {
	// 	//帳號相關
	// 	verifyToken.POST("/Logout", Controller.Logout)
	// }

	g.Run(":8080")
}

func verifyJwtToken(context *gin.Context) {
	a := context.Request.Header.Get("token")
	result, errMsg := Service.VerifyJwtToken(a, context.FullPath())
	if result {
		//呼叫下一個middleware
		context.Next()
	} else {
		context.AbortWithStatusJSON(200, CommonModel.ApiResponse{
			Success: false,
			Code:    ResultCode.InvalidToken,
			Message: "Valid faliure. " + errMsg,
		})
	}
}

func PanicHandler(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			context.AbortWithStatusJSON(200, CommonModel.ApiResponse{
				Success: false,
				Code:    ResultCode.Unknown,
				Message: fmt.Sprint(r),
			})
		}
	}()

	//呼叫下一個middleware
	context.Next()
}

/*
因為GIN無法直接讀取Response Body，
所以自己建一個Writer給GIN用，這個新的Writer會在Gin每次呼叫Write的時候，將Write的內容多寫一份至customWriter.body。
*/
type customWriter struct {
	gin.ResponseWriter
	customBody *bytes.Buffer
}

func (w customWriter) Write(b []byte) (int, error) {
	w.customBody.Write(b)
	return w.ResponseWriter.Write(b)
}

func CustomLogger(ginContext *gin.Context) {
	guid := ginContext.Request.Header.Get(ConfigHelper.RequestGuidKey)
	if guid == "" {
		guid = uuid.New().String()
	}
	ConfigHelper.SetString(ConfigHelper.RequestGuidKey, guid)

	startTime := TimeHelper.GetUTC8TimeNow()
	//讀取Request Body
	requestBody, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		LogHelper.LogInformation("CostumLogger Error At Read Request Body.")
	}

	//設定Log內容
	log := CommonModel.CommonLogContent{
		LogId:       guid,
		Route:       ginContext.Request.RequestURI,
		HttpMethod:  ginContext.Request.Method,
		Header:      ginContext.Request.Header,
		Request:     string(requestBody),
		RequestTime: TimeHelper.TimeToStringWithMiliSecond(startTime),
		LogTimeZone: "UTC+8",
		Type:        "GinLogger",
	}

	//重置request body reader，讓controller能正常讀取Body的內容
	ginContext.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	//因為無法直接讀寫gin的ResponseBody，將原生的gin.writer加一層包裝，讓gin.writer.Write的時候，會多寫一份到customWriter.customBody
	customWriter := customWriter{
		customBody:     bytes.NewBufferString(""),
		ResponseWriter: ginContext.Writer,
	}
	ginContext.Writer = customWriter

	ginContext.Next()

	//補上Response的Log
	endTime := TimeHelper.GetUTC8TimeNow()
	log.Response = customWriter.customBody.String()
	log.ResponseTime = TimeHelper.TimeToStringWithMiliSecond(endTime)
	log.ExecuteSecond = endTime.Sub(startTime).Seconds()
	log.HttpStatusCode = ginContext.Writer.Status()
	LogHelper.LogInformation(StructHelper.ToJsonString(log))
}
