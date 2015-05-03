package helper

import (
	"errors"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/tango-contrib/renders"
	"path/filepath"
)

const defaultTheme string = "default"

var (
	theme *ThemeManager
)

type ThemeManager struct {
	currentTheme   string
	themeDirectory string
}

func initTheme(cfg *core.Config, current string) {
	if current == "" {
		current = defaultTheme
	}
	theme = &ThemeManager{
		currentTheme:   current,
		themeDirectory: cfg.ThemeDirectory,
	}
}

// change theme caller
func setThemeCaller(v interface{}) *core.ActionResult {
	if name, ok := v.(string); ok {
		theme.currentTheme = name
		return core.NewOKActionResult(core.AData{"theme": name})
	}
	return core.NewSystemErrorResult(errors.New("theme name is not valid"))
}

// change theme
func SetTheme(name string) {
	base.Action.Call(setThemeCaller, name)
}

// theme info to view
type themeInfo struct {
	Current string
	Path    string
}

// theme controller interface
type IThemeController interface {
	Assign(key string, value interface{})
	Render(tpl string)
}

// theme controller, base on tango's render
type ThemeController struct {
	renders.Renderer
	data map[string]interface{}
}

// assign data to view
func (t *ThemeController) Assign(name string, value interface{}) {
	if len(t.data) == 0 {
		t.data = make(map[string]interface{})
		// assign theme data for default
		t.data["Theme"] = themeInfo{
			Current: theme.currentTheme,
			Path:    filepath.Join("/theme", theme.currentTheme),
		}
	}
	if name != "" {
		t.data[name] = value
	}
}

// render theme file in caller
func (t *ThemeController) Render(tpl string) {
	result := base.Action.Call(t.renderTheme, tpl)
	if !result.Meta.Status {
		panic(result.Meta.ErrorMessage)
	}
}

// render theme file
func (t *ThemeController) renderTheme(v interface{}) *core.ActionResult {
	if name, ok := v.(string); ok {
		// call assign to make sure that theme info are assigned
		if len(t.data) == 0 {
			t.Assign("", nil)
		}
		tpl := filepath.Join(theme.currentTheme, name)
		if err := t.Renderer.Render(tpl, t.data); err != nil {
			return core.NewSystemErrorResult(err)
		}
		return core.NewOKActionResult(core.AData{
			"theme":    theme.currentTheme,
			"template": name,
			"data":     t.data,
		})
	}
	return core.NewSystemErrorResult(errors.New("template name is invalid"))
}
