package usershndlr

import (
	"net/http"

	"github.com/hossein1376/gotp/pkg/domain/model"
	"github.com/hossein1376/gotp/pkg/handler/rest/serde"
	"github.com/hossein1376/gotp/pkg/service/usersrvc"
)

type UsersHandler struct {
	userService *usersrvc.UserService
}

func NewUsersHandler(userService *usersrvc.UserService) *UsersHandler {
	return &UsersHandler{userService: userService}
}

func (h UsersHandler) GetUserByPhoneHandler(w http.ResponseWriter, r *http.Request) {
	phone := r.PathValue("phone")
	if phone == "" || !model.PhoneRegEx.MatchString(phone) {
		resp := serde.Response{Message: "invalid phone number"}
		serde.WriteJson(w, http.StatusBadRequest, resp, nil)
		return
	}

	user, err := h.userService.GetByPhone(r.Context(), phone)
	if err != nil {
		status, resp := serde.ExtractFromErr(err)
		serde.WriteJson(w, status, resp, nil)
		return
	}

	resp := GetUserByPhoneResponse{User: user}
	serde.WriteJson(w, http.StatusOK, resp, nil)
	return
}

func (h UsersHandler) ListAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	req, err := bindListUsersRequest(r)
	if err != nil {
		resp := serde.Response{Message: err.Error()}
		serde.WriteJson(w, http.StatusBadRequest, resp, nil)
		return
	}

	users, err := h.userService.ListUsers(
		r.Context(), req.count, req.page, req.offset, req.desc,
	)
	if err != nil {
		status, resp := serde.ExtractFromErr(err)
		serde.WriteJson(w, status, resp, nil)
		return
	}

	resp := ListUsersResponse{Users: users}
	serde.WriteJson(w, http.StatusOK, resp, nil)
	return
}
