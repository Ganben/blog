package base

import (
	"github.com/gofxh/blog/app"
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
	"html/template"
	"path"
)

var (
	_ ThemeRoute       = new(PageRouter)
	_ ThemeRenderRoute = new(PageRouter)
)

type PageRouter struct {
	tango.Ctx
	renders.Renderer

	viewData map[string]interface{}
	viewFunc template.FuncMap
}

// assign view data
func (pr *PageRouter) Assign(key string, value interface{}) {
	if len(pr.viewData) == 0 {
		pr.viewData = make(map[string]interface{})
		pr.viewData["ThemeLink"] = pr.GetThemeLink()
	}
	pr.viewData[key] = value
}

// assign view function
/*
func (pr *PageRouter) AssignFunc(key string, fn interface{}) {
	if len(pr.viewFunc) == 0 {
		pr.viewFunc = make(template.FuncMap)
	}
	pr.viewFunc[key] = fn
}*/

// render theme file to bytes
func (pr *PageRouter) RenderThemeBytes(file string) ([]byte, error) {
	return pr.RenderBytes(pr.GetThemeFile(file), pr.viewData)
}

// render theme file to response
func (pr *PageRouter) RenderTheme(status int, file string) error {
	return pr.StatusRender(status, pr.GetThemeFile(file), pr.viewData)
}

// must render theme to response, otherwise panic
func (pr *PageRouter) MustRenderTheme(status int, file string) {
	if err := pr.RenderTheme(status, file); err != nil {
		panic(err)
	}
}

// get current theme name
func (pr *PageRouter) GetTheme() string {
	return app.Theme.GetName()
}

// get current theme path
func (pr *PageRouter) GetThemePath() string {
	return app.Theme.GetPath()
}

// get current theme url
func (pr *PageRouter) GetThemeLink() string {
	return app.Theme.GetLink()
}

// get current theme file
func (pr *PageRouter) GetThemeFile(file string) string {
	return path.Join(app.Theme.GetName(), file)
}

// get current theme file path
func (pr *PageRouter) GetThemeFilePath(file string) string {
	return path.Join(app.Theme.GetPath(), file)
}

// get current theme file url
func (pr *PageRouter) GetThemeFileLink(file string) string {
	return path.Join(app.Theme.GetLink(), file)
}

type AdminPageRouter struct {
	tango.Ctx
	renders.Renderer

	viewData map[string]interface{}
	viewFunc template.FuncMap
}

// assign view data
func (pr *AdminPageRouter) Assign(key string, value interface{}) {
	if len(pr.viewData) == 0 {
		pr.viewData = make(map[string]interface{})
		pr.viewData["ThemeLink"] = pr.GetThemeLink()
	}
	pr.viewData[key] = value
}

// assign view function
/*
func (pr *PageRouter) AssignFunc(key string, fn interface{}) {
	if len(pr.viewFunc) == 0 {
		pr.viewFunc = make(template.FuncMap)
	}
	pr.viewFunc[key] = fn
}*/

// render admin theme file to bytes
func (pr *AdminPageRouter) RenderThemeBytes(file string) ([]byte, error) {
	return pr.RenderBytes(pr.GetThemeFile(file), pr.viewData)
}

// render admin theme file to response
func (pr *AdminPageRouter) RenderTheme(status int, file string) error {
	return pr.StatusRender(status, pr.GetThemeFile(file), pr.viewData)
}

// must render admin theme to response, otherwise panic
func (pr *AdminPageRouter) MustRenderTheme(status int, file string) {
	if err := pr.RenderTheme(status, file); err != nil {
		panic(err)
	}
}

// get admin theme name
func (pr *AdminPageRouter) GetTheme() string {
	return "admin"
}

// get admin theme path
func (pr *AdminPageRouter) GetThemePath() string {
	return app.Theme.GetAdminPath()
}

// get admin theme url
func (pr *AdminPageRouter) GetThemeLink() string {
	return app.Theme.GetAdminLink()
}

// get admin theme file
func (pr *AdminPageRouter) GetThemeFile(file string) string {
	return path.Join("admin", file)
}

// get admin theme file path
func (pr *AdminPageRouter) GetThemeFilePath(file string) string {
	return path.Join(app.Theme.GetAdminPath(), file)
}

// get admin theme file url
func (pr *AdminPageRouter) GetThemeFileLink(file string) string {
	return path.Join(app.Theme.GetAdminLink(), file)
}
