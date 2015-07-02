package action

import (
	"errors"
	"github.com/gofxh/blog/app/model"
)

var (
	ERR_USERNAME_UNIQUE   = errors.New("username-need-unique")
	ERR_USER_EMAIL_UNIQUE = errors.New("user-email-need-unique")
	ERR_PASSWORD_CONFIRM  = errors.New("password-confirm-error")
)

type ProfileForm struct {
	Id    int64  `form:"id" binding:"Required"`
	Name  string `form:"username" binding:"Required:AlphaDash"`
	Nick  string `form:"nick" binding:"Required"`
	Email string `form:"email" binding:"Required;Email"`
	Url   string `form:"url" binding:"Url"`
	Bio   string `form:"bio"`
}

// update user profile
func UserUpdateProfile(v interface{}) *Result {
	form, ok := v.(*ProfileForm)
	if !ok {
		return ErrorResult(paramTypeError(new(ProfileForm)))
	}

	// check name unique
	u, err := model.GetUserByUniqueName(form.Id, form.Name)
	if err != nil {
		return ErrorResult(err)
	}
	if u != nil && u.Id != form.Id {
		return ErrorResult(ERR_USERNAME_UNIQUE)
	}

	// check email
	u, err = model.GetUserByUniqueEmail(form.Id, form.Email)
	if err != nil {
		return ErrorResult(err)
	}
	if u != nil && u.Id != form.Id {
		return ErrorResult(ERR_USER_EMAIL_UNIQUE)
	}

	// update user
	user := &model.User{
		Id:    form.Id,
		Name:  form.Name,
		Nick:  form.Nick,
		Email: form.Email,
		Url:   form.Url,
		Bio:   form.Bio,
	}
	u, err = model.UpdateUserProfile(user)
	if err != nil {
		return ErrorResult(err)
	}

	return OkResult(map[string]interface{}{
		"user": u,
	})
}

type PasswordForm struct {
	Id      int64  `form:"id" binding:"Required"`
	Old     string `form:"old-password" binding:"Required;AlphaDashDot"`
	New     string `form:"new-password" binding:"Required;AlphaDashDot"`
	Confirm string `form:"confirm-password" binding:"Required;AlphaDashDot"`
}

// update password
func UserUpdatePassword(v interface{}) *Result {
	form, ok := v.(*PasswordForm)
	if !ok {
		return ErrorResult(paramTypeError(new(PasswordForm)))
	}

	// check confirm password
	if form.New != form.Confirm {
		return ErrorResult(ERR_PASSWORD_CONFIRM)
	}

	// get old user
	u, err := model.GetUserBy("id", form.Id)
	if err != nil {
		return ErrorResult(err)
	}
	// user not found
	if u == nil || u.Id != form.Id {
		return ErrorResult(ERR_USERNAME_NOT_FOUND)
	}
	// check password
	if !u.CheckPassword(form.Old) {
		return ErrorResult(ERR_PASSWORD_INCORRECT)
	}
	u.Password = form.New

	// save password
	if err = model.UpdateUserPassword(u); err != nil {
		return ErrorResult(err)
	}

	return OkResult(nil)
}
