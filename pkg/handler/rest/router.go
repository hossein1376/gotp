package rest

import (
	"net/http"

	"github.com/hossein1376/gotp/pkg/handler/rest/loginhndlr"
	"github.com/hossein1376/gotp/pkg/handler/rest/usershndlr"
	"github.com/hossein1376/gotp/pkg/service"
)

func newRouter(srvc *service.Services) *http.ServeMux {
	mux := http.NewServeMux()

	lh := loginhndlr.NewLoginHandler(srvc.LoginService)
	uh := usershndlr.NewUsersHandler(srvc.UserService)

	mux.HandleFunc("POST /api/v1/login/otp", lh.SendLoginOTPHandler)
	mux.HandleFunc("POST /api/v1/login", lh.LoginViaOTPHandler)
	mux.HandleFunc("GET /api/v1/users", uh.ListAllUsersHandler)
	mux.HandleFunc("GET /api/v1/users/{phone}", uh.GetUserByPhoneHandler)

	return mux
}
