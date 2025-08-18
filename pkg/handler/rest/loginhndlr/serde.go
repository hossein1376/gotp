package loginhndlr

type SendOTPRequest struct {
	Phone string `json:"phone"`
}

type LoginViaOTPRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type LoginViaOTPResponse struct {
	Token string `json:"token"`
}
