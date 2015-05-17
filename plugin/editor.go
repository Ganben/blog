package plugin

import (
	"github.com/Unknwon/com"
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/core"
	"github.com/gofxh/blog/lib/log"
	"github.com/gofxh/blog/mvc/helper"
	"github.com/lunny/tango"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	base.Plugin.Register(new(EditorMdPlugin))
}

// editor-md plugin
type EditorMdPlugin struct {
	enableStatus bool
}

// init plugin
func (ep *EditorMdPlugin) Init() {
	// set server static func
	base.Server.Use(tango.Static(tango.StaticOptions{
		RootPath: "plugin/editor_md",
		Prefix:   "plugin/editor_md",
	}))
	// add to action
	base.Action.After(helper.RichEditor, func(param interface{}, res *core.ActionResult) {
		// if not enable, stop it
		if !ep.IsEnabled() {
			return
		}
		// copy editor_md.html to current theme
		ep.copyFile()
		// if file valid, go next
		if ep.isValid() {
			form := param.(*helper.RichEditorForm)
			form.Template = "plugin/editor_md.html"
			bytes := form.Render.Bytes(form.Template)
			// replace with new bytes
			if len(bytes) > 0 {
				res.Data["bytes"] = bytes
			}
		}
	})
	// default enable
	ep.Enable()
}

func (ep *EditorMdPlugin) Name() string {
	return "Editor-Md Markdown Editor"
}

// is enabled
func (ep *EditorMdPlugin) IsEnabled() bool {
	return ep.enableStatus
}

// enable plugin
func (ep *EditorMdPlugin) Enable() {
	ep.enableStatus = true
}

// disable plugin
func (ep *EditorMdPlugin) Disable() {
	ep.enableStatus = false
}

// copy file
func (ep *EditorMdPlugin) copyFile() {
	// check theme path file
	file := helper.BuildThemePath("plugin/editor_md.html")
	if !com.IsFile(file) {
		bytes, err := ioutil.ReadFile("plugin/editor_md/editor_md.html")
		if err != nil {
			log.Warn("CopyFile|%s", err.Error())
			return
		}
		// create directory
		os.Mkdir(filepath.Dir(file), os.ModePerm)
		// write to template
		ioutil.WriteFile(file, bytes, os.ModePerm)
	}
}

// is file valid
func (ep *EditorMdPlugin) isValid() bool {
	return com.IsFile(helper.BuildThemePath("plugin/editor_md.html"))
}
