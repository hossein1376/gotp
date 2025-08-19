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

	mux.Handle(
		"POST /api/v1/login/otp",
		applyMiddlewares(http.HandlerFunc(lh.SendLoginOTPHandler)),
	)
	mux.Handle(
		"POST /api/v1/login",
		applyMiddlewares(http.HandlerFunc(lh.LoginViaOTPHandler)),
	)
	mux.Handle(
		"GET /api/v1/users",
		applyMiddlewaresAuth(http.HandlerFunc(uh.ListAllUsersHandler)),
	)
	mux.Handle(
		"GET /api/v1/users/{phone}",
		applyMiddlewaresAuth(http.HandlerFunc(uh.GetUserByPhoneHandler)),
	)

	return mux
}

func applyMiddlewares(h http.Handler) http.Handler {
	return recoverMiddleware(requestIDMiddleware(loggerMiddleware(h)))
}

func applyMiddlewaresAuth(h http.Handler) http.Handler {
	return applyMiddlewares(jwtAuthMiddleware(h))
}
