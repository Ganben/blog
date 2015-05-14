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
	return core.NewErrorResult(errors.New("theme name is not valid"))
}

// change theme
func SetTheme(name string) {
	base.Action.Call(setThemeCaller, name)
}

// get current theme
func GetTheme() string {
	return theme.currentTheme
}

// theme info to view
type themeInfo struct {
	Current string
	Path    string
}

// theme controller interface
type IThemeController interface {
	Assign(key string, value interface{})        // assign view data
	Render(tpl string)                           // render template, call RenderAction
	RenderAction(interface{}) *core.ActionResult // render action caller
	Bytes(tpl string) []byte                     // render template to bytes, call BytesAction
	BytesAction(interface{}) *core.ActionResult  // render bytes action caller
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
			Path:    "/theme/" + theme.currentTheme,
		}
	}
	if name != "" {
		t.data[name] = value
	}
}

// render theme file in caller
func (t *ThemeController) Render(tpl string) {
	result := base.Action.Call(t.RenderAction, tpl)
	if !result.Meta.Status {
		panic(result.Meta.ErrorMessage)
	}
}

// render theme file
func (t *ThemeController) RenderAction(v interface{}) *core.ActionResult {
	if name, ok := v.(string); ok {
		// call assign to make sure that theme info are assigned
		if len(t.data) == 0 {
			t.Assign("", nil)
		}
		tpl := filepath.Join(theme.currentTheme, name)
		if err := t.Renderer.Render(tpl, t.data); err != nil {
			return core.NewErrorResult(err)
		}
		return core.NewOKActionResult(core.AData{
			"theme":    theme.currentTheme,
			"template": name,
			"data":     t.data,
		})
	}
	return core.NewErrorResult(errors.New("template name is invalid"))
}

// render bytes in caller
func (t *ThemeController) Bytes(tpl string) []byte {
	result := base.Action.Call(t.BytesAction, tpl)
	if !result.Meta.Status {
		panic(result.Meta.ErrorMessage)
	}
	return result.Data["bytes"].([]byte)
}

// render theme to bytes
func (t *ThemeController) BytesAction(v interface{}) *core.ActionResult {
	if name, ok := v.(string); ok {
		// call assign to make sure that theme info are assigned
		if len(t.data) == 0 {
			t.Assign("", nil)
		}
		tpl := filepath.Join(theme.currentTheme, name)
		bytes, err := t.Renderer.RenderBytes(tpl, t.data)
		if err != nil {
			return core.NewErrorResult(err)
		}
		return core.NewOKActionResult(core.AData{
			"theme":    theme.currentTheme,
			"template": name,
			"data":     t.data,
			"bytes":    bytes,
		})
	}
	return core.NewErrorResult(errors.New("template name is invalid"))
}

var (
	RICH_EDITOR_BAD_DATA = errors.New("bad data")
)

type RichEditorForm struct {
	Render   ThemeController
	Template string
}

// rich editor
func RichEditor(v interface{}) *core.ActionResult {
	form, ok := v.(*RichEditorForm)
	if !ok {
		return core.NewErrorResult(RICH_EDITOR_BAD_DATA)
	}
	bytes := form.Render.Bytes(form.Template)
	return core.NewOKActionResult(core.AData{
		"bytes": bytes,
	})
}
