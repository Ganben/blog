package core

import (
	"github.com/Unknwon/com"
	"os"
	"path"
	"strings"
)

// global theme manager
type Theme struct {
	directory string
	name      string
}

// new theme manager
func NewTheme(dir, name string) *Theme {
	// try to create dir
	if !com.IsDir(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return &Theme{dir, name}
}

// get theme link, for http server
func (t *Theme) GetLink() string {
	return path.Join("/theme", strings.ToLower(t.name))
}

// get theme path
func (t *Theme) GetPath() string {
	return path.Join(t.directory, strings.ToLower(t.name))
}

// get theme name
func (t *Theme) GetName() string {
	return t.name
}

// set theme name
func (t *Theme) SetName(name string) {
	t.name = name
}

// get admin theme link, for http server
func (t *Theme) GetAdminLink() string {
	return "/theme/admin"
}

// get admin theme path
func (t *Theme) GetAdminPath() string {
	return path.Join(t.directory, "admin")
}
