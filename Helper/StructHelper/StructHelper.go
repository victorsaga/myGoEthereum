package StructHelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/ahmetb/go-linq/v3"
)

func ToStruct(jsonString string, v interface{}) {
	err := json.Unmarshal([]byte(jsonString), &v)
	if err != nil {
		panic(err)
	}
}

func ToJsonString(inputStruct interface{}) (outPut string) {
	inputStructBytes, _ := json.Marshal(inputStruct)

	return string(inputStructBytes)
}

func ToMap(inputStruct interface{}) (outPut map[string]interface{}) {
	inputStructBytes, _ := json.Marshal(inputStruct)

	json.Unmarshal(inputStructBytes, &outPut)

	return
}

func ToQueryString(inputStruct interface{}, ignore []string) string {
	if inputStruct == nil {
		return ""
	}
	return MapToQueryString(ToMap(inputStruct), ignore)
}

func myFunc() {
	fmt.Print(strconv.Itoa(1) + strconv.Itoa(int(time.Second)))
}

func MapToQueryString(value map[string]interface{}, ignore []string) (result string) {
	if value == nil || len(value) == 0 {
		return
	}

	i := linq.From(ignore)
	for k, v := range value {
		if i.AnyWithT(func(s string) bool {
			return s == k
		}) {
			continue
		}

		if len(result) > 0 {
			result += "&"
		}
		result += k + "=" + ToString(v)
	}
	return
}

func ToString(arg interface{}, timeFormat ...string) string {
	if len(timeFormat) > 1 {
		log.SetFlags(log.Llongfile | log.LstdFlags)
		log.Println(errors.New(fmt.Sprintf("timeFormat's length should be one")))
	}
	var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		if len(timeFormat) == 1 {
			return v.Format(timeFormat[0])
		}
		return v.Format("2006-01-02 15:04:05")
	case fmt.Stringer:
		return v.String()
	case reflect.Value:
		return ToString(v.Interface(), timeFormat...)
	default:
		return ""
	}
}
