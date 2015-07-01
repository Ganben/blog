package action

import (
	"errors"
	"fmt"
	"github.com/gofxh/blog/app/model"
	"time"
)

var (
	ERR_USERNAME_NOT_FOUND = errors.New("username-not-found")
	ERR_PASSWORD_INCORRECT = errors.New("password-incorrect")
)

// login form
type LoginForm struct {
	Username string `form:"username" binding:"Required"`
	Password string `form:"password" binding:"Required;AlphaDash"`
	Remember int    `form:"remember"`
}

func UserLogin(v interface{}) *Result {
	form, ok := v.(*LoginForm)
	if !ok {
		return ErrorResult(paramTypeError(new(LoginForm)))
	}
	// get user
	u, err := model.GetUserBy("name", form.Username)
	if err != nil {
		return ErrorResult(err)
	}
	if u.Id == 0 {
		// user not found
		return ErrorResult(ERR_USERNAME_NOT_FOUND)
	}
	// check password
	if !u.CheckPassword(form.Password) {
		return ErrorResult(ERR_PASSWORD_INCORRECT)
	}

	// create token
	t := &model.Token{
		UserId:     u.Id,
		Value:      fmt.Sprintf("%d", time.Now().UnixNano()),
		ExpireTime: time.Now().Add(24 * time.Hour).Unix(),
	}
	t.EncodeValue()
	if err = model.SaveToken(t); err != nil {
		return ErrorResult(err)
	}

	// return user and token data
	return OkResult(map[string]interface{}{
		"user":  u,
		"token": t,
	})
}
