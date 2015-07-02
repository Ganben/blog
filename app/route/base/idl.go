package base

import (
	"github.com/gofxh/blog/app/model"
	"github.com/lunny/tango"
)

type ThemeRoute interface {
	GetTheme() string                    // get current theme name
	GetThemePath() string                // get theme directory path
	GetThemeLink() string                // get theme base url
	GetThemeFile(file string) string     // get theme file base on theme directory
	GetThemeFilePath(file string) string // get theme file path
	GetThemeFileLink(file string) string // get theme file link url
}

type ThemeRenderRoute interface {
	ThemeRoute
	Assign(key string, value interface{}) // assign view data
	// AssignFunc(key string, fn interface{})        // assign template func, seems no useful
	RenderThemeBytes(file string) ([]byte, error) // render theme to bytes
	RenderTheme(status int, file string) error    // render
	MustRenderTheme(status int, file string)      // must render and panic
}

type BindRoute interface {
	BindAndValidate(interface{}) error // bind and validate, then return error message
}

type AuthRoute interface {
	GetAuthToken(ctx *tango.Context) string // get token for authorization
	SetAuthUser(*model.User)                // set authorized user
	GetAuthSuccessRedirect() string         // get redirect url if authorized
	GetAuthFailRedirect() string            // get redirect url if not authorized
}
