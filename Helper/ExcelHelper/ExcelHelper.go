package ExcelHelper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/tealeg/xlsx"
)

//只收陣列
//可以使用Tag ExcelHelper:"-"，避免將某些欄位轉換成excel
func ToXlsx(input interface{}) (result io.ReadSeeker, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	//驗證input是否為陣列
	inputKind := reflect.TypeOf(input).Kind()
	if inputKind != reflect.Slice {
		err = errors.New("Input is not array.")
		return
	}

	//取得物件欄位名稱
	t := reflect.TypeOf(input).Elem()
	fields := []string{}
	fieldsShowName := []string{}
	var fieldShowName string
	for i := 0; i < t.NumField(); i++ {
		//如果有ExcelHelper的TAG，且值為-，就不將此欄位列入取值名單
		tag := t.Field(i).Tag.Get("ExcelHelper")

		if tag == "-" {
			continue
		}

		if len(tag) > 0 {
			fieldShowName = tag
		} else {
			fieldShowName = t.Field(i).Name
		}
		fieldsShowName = append(fieldsShowName, fieldShowName)
		fields = append(fields, t.Field(i).Name)
	}

	//產生試算表
	file := xlsx.NewFile()

	//加試算表分頁
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	// 插入標頭
	titleRow := sheet.AddRow()
	for _, v := range fieldsShowName {
		cell := titleRow.AddCell()
		cell.Value = v
	}

	//插入資料
	s := reflect.ValueOf(input)

	for i := 0; i < s.Len(); i++ {
		row := sheet.AddRow()

		for _, v := range fields {
			cell := row.AddCell()

			f := reflect.Indirect(s.Index(i)).FieldByName(v)

			//處理指標
			if f.Kind() == reflect.Pointer {
				if f.Elem().CanAddr() {
					cell.SetValue(f.Elem().Interface())
				} else {
					cell.SetValue("")
				}
			} else {
				cell.SetValue(f.Interface())
			}
		}
	}

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	result = bytes.NewReader(buffer.Bytes())
	return
}
