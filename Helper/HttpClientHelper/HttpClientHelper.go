package HttpClientHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"myGoEthereum/Helper/LogHelper"
	"myGoEthereum/Helper/StructHelper"
	"myGoEthereum/Helper/TimeHelper"
	"myGoEthereum/Model/CommonModel"
	"net/http"
	"strings"
)

type ContentType string
type Method string

const (
	MethodGet  Method = http.MethodGet
	MethodPost Method = http.MethodPost

	ContentTypeForm ContentType = "application/x-www-form-urlencoded"
	ContentTypeJSON ContentType = "application/json"
)

func Send(
	myUrl string,
	method Method,
	contentType ContentType,
	header interface{},
	body interface{},
	logId string) (httpStatusCode int, respBody string, err error) {
	startTime := TimeHelper.GetUTC8TimeNow()
	log := CommonModel.CommonLogContent{
		LogId:       logId,
		Route:       myUrl,
		HttpMethod:  string(method),
		RequestTime: TimeHelper.TimeToStringWithMiliSecond(startTime),
		LogTimeZone: "UTC+8",
		Type:        "HttpClientHelperLogger",
	}

	//錯誤處理
	defer func() {
		r := recover()
		if r != nil {
			log.Error = fmt.Sprint(r)
			LogHelper.LogInformation(StructHelper.ToJsonString(log))
			err = r.(error)
		}
	}()

	//處理body
	var bb []byte
	var ioReader io.Reader
	var queryString string
	if method == MethodGet {
		queryString = StructHelper.ToQueryString(body, nil)
		log.Request = queryString
		if len(queryString) > 0 {
			myUrl += "?"
		}
		myUrl += queryString
	} else {
		switch contentType {
		case ContentTypeJSON:
			bb, _ = json.Marshal(&body)
			ioReader = bytes.NewReader(bb)
			log.Request = string(bb)
		case ContentTypeForm:
			queryString = StructHelper.ToQueryString(body, nil)
			ioReader = strings.NewReader(queryString)
			log.Request = queryString
		}
	}
	//建立請求參數物件
	req, err := http.NewRequest(string(method), myUrl, ioReader)
	if err != nil {
		panic(err)
	}

	//處理header
	req.Header.Add("Content-Type", string(contentType))
	if header != nil {
		m := StructHelper.ToMap(header)

		for field, val := range m {
			req.Header.Add(field, fmt.Sprint(val))
		}
	}
	log.Header = fmt.Sprint(req.Header)

	//送出
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	endTime := TimeHelper.GetUTC8TimeNow()
	log.ResponseTime = TimeHelper.TimeToStringWithMiliSecond(endTime)
	log.ExecuteSecond = endTime.Sub(startTime).Seconds()
	log.HttpStatusCode = resp.StatusCode

	httpStatusCode = resp.StatusCode

	defer resp.Body.Close()
	responseBodyByte, _ := ioutil.ReadAll(resp.Body)
	respBody = string(responseBodyByte)
	log.Response = respBody

	LogHelper.LogInformation(StructHelper.ToJsonString(log))
	return
}
