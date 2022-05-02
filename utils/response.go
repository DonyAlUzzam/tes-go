package utils

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
	Token     string      `json:"-"`
}

type ResponseLogin struct {
	ErrorCode int    `json:"error_code" form:"error_code"`
	Message   string `json:"message" form:"message"`
	Token     string `json:"token"`
}
