package core

import (
	"encoding/json"
	"github.com/Unknwon/com"
	"io/ioutil"
	"os"
)

const CONFIG_FILE = "config.json"

// config struct
type Config struct {
	AppName        string // app name
	AppVersion     string // app version
	AppInstallTime int64  // install time, unix stamp

	HttpHost    string // http server host
	HttpAddress string // http server port

	UserDirectory       string // user directory
	UserThemeDirectory  string // user theme directory, under user directory
	UserUploadDirectory string // user upload directory, under user directory
	UserDataFile        string // user sql data file
}

// new config struct,
// set default, load from json file or write json file when first run.
func NewConfig() *Config {
	c := &Config{
		AppName:             APP_NAME,
		AppVersion:          APP_VERSION,
		AppInstallTime:      0,
		HttpHost:            "0.0.0.0",
		HttpAddress:         "3098",
		UserDirectory:       "data",
		UserThemeDirectory:  "theme",
		UserUploadDirectory: "upload",
		UserDataFile:        "data.db",
	}
	// if config file exist, read file
	if com.IsFile(CONFIG_FILE) {
		bytes, _ := ioutil.ReadFile(CONFIG_FILE)
		json.Unmarshal(bytes, c)
		return c
	}
	// write config file if not exist
	c.Write()
	return c
}

// write config to file
func (c *Config) Write() {
	bytes, _ := json.MarshalIndent(c, "", "  ")
	ioutil.WriteFile(CONFIG_FILE, bytes, os.ModePerm)
}
