package helper

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/mvc/action"
	"github.com/lunny/tango"
)

// user auth controller interface
type IAuthController interface {
	SetAuthUser(*entity.User)
	GetAuthSuccessRedirect() string // auth success redirect
	GetAuthFailRedirect() string    // auth fail redirect
}

// default user auth controller
type AuthController struct {
	AuthUser *entity.User
}

// implement user auth controller
func (ac *AuthController) SetAuthUser(u *entity.User) {
	ac.AuthUser = u
}

// implement user auth controller
func (ac *AuthController) GetAuthSuccessRedirect() string {
	return ""
}

// implement user auth controller
func (ac *AuthController) GetAuthFailRedirect() string {
	return ""
}

// auth handler in tango middleware
func UseAuth() tango.HandlerFunc {
	return func(ctx *tango.Context) {
		if act, ok := ctx.Action().(IAuthController); ok {
			tokenC := ctx.Cookies().Get("token")
			if tokenC != nil {
				tokenValue := tokenC.Value
				if tokenValue != "" {
					if result := base.Action.Call(action.Auth, tokenValue); result.Meta.Status {
						act.SetAuthUser(result.Data["user"].(*entity.User))
						ctx.Next()
						return
					}
				}
			}
			if redirect := act.GetAuthFailRedirect(); redirect != "" {
				ctx.Redirect(redirect, 302)
				return
			}
		}
		ctx.Next()
	}
}
