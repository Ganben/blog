package base

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
