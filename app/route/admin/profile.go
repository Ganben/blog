package admin

import "github.com/gofxh/blog/app/route/base"

type Profile struct {
	base.AdminPageRouter
	base.AuthRouter
}

func (p *Profile) Get() {
	p.Assign("Title", "Profile")
	p.Assign("IsProfilePage", true)
	p.Assign("AuthUser", p.AuthUser)
	p.MustRenderTheme(200, "profile.tmpl")
}
