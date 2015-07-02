package admin

import (
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/model"
	"github.com/gofxh/blog/app/route/base"
	"github.com/lunny/tango"
)

// profile route
type Profile struct {
	base.AdminPageRouter
	base.AuthRouter
	base.BindRouter
}

// show profile page
func (p *Profile) Get() {
	pwdMessage := p.Form("password")
	if pwdMessage != "" {
		if pwdMessage == "success" {
			p.Assign("PasswordSuccess", true)
		} else {
			p.Assign("PasswordError", pwdMessage)
		}
	}
	p.Assign("Title", "Profile")
	p.Assign("IsProfilePage", true)
	p.Assign("AuthUser", p.AuthUser)
	p.MustRenderTheme(200, "profile.tmpl")
}

// update post page
func (p *Profile) Post() {
	// bind form
	form := &action.ProfileForm{}
	if err := p.BindAndValidate(form); err != nil {
		p.Assign("ProfileError", err.Error())
		p.Get()
		return
	}

	// call update action
	res := action.Call(action.UserUpdateProfile, form)
	if !res.Status {
		p.Assign("ProfileError", res.Error)
		p.Get()
		return
	}

	// update auth user data
	p.SetAuthUser(res.Data["user"].(*model.User))
	p.Assign("ProfileSuccess", true)
	p.Get()
}

// password route
type Password struct {
	tango.Ctx
	base.AuthRouter
	base.BindRouter
}

// update password post route
func (p *Password) Post() {
	form := &action.PasswordForm{}
	if err := p.BindAndValidate(form); err != nil {
		p.Redirect("/admin/profile?password=" + err.Error())
		return
	}

	// call update password action
	res := action.Call(action.UserUpdatePassword, form)
	if !res.Status {
		p.Redirect("/admin/profile?password=" + res.Error)
		return
	}

	// update auth user data
	p.Redirect("/admin/profile?password=true")
}
