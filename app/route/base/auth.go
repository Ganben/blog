package base

import (
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/model"
	"github.com/lunny/tango"
)

// base auth router
type AuthRouter struct {
	AuthUser *model.User
}

// get auth token via http header, cookie or form value
func (ar *AuthRouter) GetAuthToken(ctx *tango.Context) string {
	var token string
	if token = ctx.Header().Get("X-Token"); token != "" {
		return token
	}
	if token = ctx.Cookie("x-token"); token != "" {
		return token
	}
	return ctx.Form("x-token")
}

// set auth user
func (ar *AuthRouter) SetAuthUser(u *model.User) {
	ar.AuthUser = u
}

// get success redirect url
func (ar *AuthRouter) GetAuthSuccessRedirect() string {
	return ""
}

// get fail redirect url
func (ar *AuthRouter) GetAuthFailRedirect() string {
	return "/admin/logout"
}

// auth handler
func AuthHandler() tango.HandlerFunc {
	return func(ctx *tango.Context) {
		auth, ok := ctx.Action().(AuthRoute)
		if !ok {
			ctx.Next()
			return
		}
		// read token
		token := auth.GetAuthToken(ctx)
		if token != "" {
			result := action.Call(action.UserAuth, token)
			if result.Status {
				auth.SetAuthUser(result.Data["user"].(*model.User))
				ctx.Next()
				return
			}
		}
		// fail redirect
		if url := auth.GetAuthFailRedirect(); url != "" {
			ctx.Redirect(url, 302)
			return
		}

		// auth fail , no redirect, to show 401
		ctx.WriteHeader(401)
	}
}
