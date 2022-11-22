package BaseModel

import (
	"myGoEthereum/Model/CommonModel"

	"github.com/golang-jwt/jwt"
)

const ExchangeTypeNonHedging int8 = 0
const ExchangeTypeHedging int8 = 1

type RequestPage struct {
	Page     *int `json:"page"`
	PageSize *int `json:"pageSize"`
}
type ResponsePage struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `db:"total" json:"total"`
}

type LoginRequest struct {
	Account  string
	Password string
}

type LoginResponse struct {
	CommonModel.ApiResponse
	Data LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}

type JwtPayload struct {
	jwt.StandardClaims
	FunctionIds []int
	Account     string
	Expires     int64
}
