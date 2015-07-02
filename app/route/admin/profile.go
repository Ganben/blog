package admin

import (
	"github.com/gofxh/blog/app/action"
	"github.com/gofxh/blog/app/model"
	"github.com/gofxh/blog/app/route/base"
)

type Profile struct {
	base.AdminPageRouter
	base.AuthRouter
	base.BindRouter
}

func (p *Profile) Get() {
	p.Assign("Title", "Profile")
	p.Assign("IsProfilePage", true)
	p.Assign("AuthUser", p.AuthUser)
	p.MustRenderTheme(200, "profile.tmpl")
}

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
