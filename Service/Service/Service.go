package Service

import (
	"crypto/md5"
	"fmt"
	"myGoEthereum/Helper/ConfigHelper"
	"myGoEthereum/Model/BaseModel"
	"myGoEthereum/Model/ResultCode"
	"myGoEthereum/Repository/Repository"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(r BaseModel.LoginRequest) (response BaseModel.LoginResponse) {
	if r.Account == "" || r.Password == "" {
		response.SetError(ResultCode.Parameter, "Missing Parameter.")
		return
	}

	//從資料庫取出密碼
	hashPwd := Repository.GetAccountPassword(r.Account)

	if hashPwd == nil {
		response.SetError(ResultCode.Parameter, "Login Failed, Account or Password Invalid.")
		return
	}

	//如果密碼不正確
	if !strings.EqualFold(*hashPwd, doMd5(r.Password)) {
		response.SetError(ResultCode.AccountNameIsAlreadyUsed, "Login Failed, Account or Password Invalid.")
		return
	}
	response.Data.AccessToken = createJwtToken(r.Account, []int{})

	return
}

func doMd5(input string) (output string) {
	data := []byte(input)
	has := md5.Sum(data)
	output = fmt.Sprintf("%x", has) //将[]byte转成16进制
	return
}

func createJwtToken(accountName string, functionIds []int) (output string) {
	jwtKey := []byte(ConfigHelper.GetString("JwtSettings.SignKey"))

	payload := BaseModel.JwtPayload{
		Account:     accountName,
		FunctionIds: functionIds,
		//+8時區
		Expires: time.Now().Add(time.Hour * time.Duration(ConfigHelper.GetInt64("JwtSettings.ExpireHour"))).UnixMilli(), //7天後jwt過期
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	output, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return
}

func VerifyJwtToken(jwtToken string, route string) (bool, string) {
	if jwtToken == "" {
		return false, "Token is empty."
	}
	var payload BaseModel.JwtPayload
	jwtKey := []byte(ConfigHelper.GetString("JwtSettings.SignKey"))

	token, err := jwt.ParseWithClaims(jwtToken, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return jwtKey, nil
	})

	if err != nil {
		panic(err)
	}

	if !token.Valid {
		return false, "Token Invalid."
	}

	//檢查時間是否過期
	if payload.Expires < time.Now().UnixMilli() {
		return false, "Token Invalid."
	}

	//檢查Token是否已登出
	// if Repository.IsJwtTokenLogout(jwtToken) {
	// 	return false, "Token Invalid."
	// }

	// //檢查權限足夠
	// routeFunctionId := Repository.GetRouteFunctionId(route)
	// if routeFunctionId == nil {
	// 	return false, "Route Premission Not Setting."
	// }

	// havePremission := false
	// for _, a := range payload.FunctionIds {
	// 	if a == *routeFunctionId {
	// 		havePremission = true
	// 		break
	// 	}
	// }

	// if !havePremission {
	// 	return false, "No Premission."
	// }

	return true, ""
}
