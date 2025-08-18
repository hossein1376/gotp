package loginhndlr

import (
	"net/http"

	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/handler/rest/serde"
	"github.com/hossein1376/gotp/pkg/service/loginsrvc"
)

type LoginHandler struct {
	loginService *loginsrvc.LoginService
}

func NewLoginHandler(loginService *loginsrvc.LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

func (h LoginHandler) SendLoginOTPHandler(w http.ResponseWriter, r *http.Request) {
	req := SendOTPRequest{}
	err := serde.ReadJson(r, &req)
	if err != nil {
		resp := serde.Response{Message: err.Error()}
		serde.WriteJson(w, http.StatusBadRequest, resp, nil)
		return
	}

	if ok := model.PhoneRegEx.MatchString(req.Phone); !ok {
		resp := serde.Response{Message: "invalid phone number"}
		serde.WriteJson(w, http.StatusBadRequest, resp, nil)
		return
	}

	err = h.loginService.SendLoginOTP(r.Context(), req.Phone)
	if err != nil {
		status, resp := serde.ExtractFromErr(err)
		serde.WriteJson(w, status, resp, nil)
		return
	}

	serde.WriteJson(w, http.StatusNoContent, nil, nil)
	return
}

func (h LoginHandler) LoginViaOTPHandler(w http.ResponseWriter, r *http.Request) {
	req := LoginViaOTPRequest{}
	err := serde.ReadJson(r, &req)
	if err != nil {
		resp := serde.Response{Message: err.Error()}
		serde.WriteJson(w, http.StatusBadRequest, resp, nil)
		return
	}

	if ok := model.PhoneRegEx.MatchString(req.Phone); !ok {
		resp := serde.Response{Message: "invalid phone number"}
		serde.WriteJson(w, http.StatusBadRequest, resp, nil)
		return
	}

	err = h.loginService.LoginOTP(r.Context(), req.Phone, req.Code)
	if err != nil {
		status, resp := serde.ExtractFromErr(err)
		serde.WriteJson(w, status, resp, nil)
		return
	}

	resp := LoginViaOTPResponse{
		Token: "", //TODO
	}

	serde.WriteJson(w, http.StatusCreated, resp, nil)
	return
}
