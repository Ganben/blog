package action

import (
	"errors"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/mvc/model"
)

var (
	LOGIN_BAD_DATA_ERROR  = errors.New("bad data")
	LOGIN_NO_USER_ERROR   = errors.New("no user")
	LOGIN_WRONG_PWD_ERROR = errors.New("wrong password")

	AUTH_BAD_DATA_ERROR = errors.New("bad data")
	AUTH_NO_USER_ERROR  = errors.New("no user")
)

// login form
type LoginForm struct {
	User          string
	Password      string
	Ip, UserAgent string
	Expire        int64
}

// login action, with login form
func Login(v interface{}) *core.ActionResult {
	f, ok := v.(*LoginForm)
	if !ok {
		return core.NewErrorResult(LOGIN_BAD_DATA_ERROR)
	}
	user := model.GetUserByName(f.User)
	if user == nil {
		return core.NewErrorResult(LOGIN_NO_USER_ERROR)
	}
	if !entity.CompareUserPassword(user, f.Password) {
		return core.NewErrorResult(LOGIN_WRONG_PWD_ERROR)
	}
	token := model.CreateToken(user.Id, f.Ip, f.UserAgent, f.Expire)
	return core.NewOKActionResult(core.AData{
		"user":  user,
		"token": token,
	})
}

// auth action , with token value
func Auth(v interface{}) *core.ActionResult {
	str, ok := v.(string)
	if !ok {
		return core.NewErrorResult(AUTH_BAD_DATA_ERROR)
	}
	token, user := model.GetUserByTokenValue(str)
	// extend this token
	model.ExtendToken(token)
	if user == nil {
		return core.NewErrorResult(AUTH_NO_USER_ERROR)
	}
	return core.NewOKActionResult(core.AData{
		"user": user,
	})
}
