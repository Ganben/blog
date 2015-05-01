package core

import (
	"encoding/json"
	"github.com/Unknwon/com"
	"github.com/gofxh/blog/lib/log"
	"io/ioutil"
	"os"
	"time"
)

var configFile string = "config.json" // default json data

// config data
type Config struct {
	InstallTime    int64  `json:"install_time"`
	AppVersion     string `json:"app_version"`
	AppVersionDate string `json:"app_version_date"`

	HttpAddress   string `json:"http_address"`
	DataDirectory string `json:"data_directory"`
}

// new config data,
// it loads config file if exist,
// or its fields are filled by default values
func NewConfig() *Config {
	c := new(Config)

	// read from file if exist
	if com.IsFile(configFile) {
		if err := c.ReadFile(); err != nil {
			log.Error("Config|ReadFile|%s", err.Error())
		}
	}

	// fill default if config is not valid
	if c.AppVersion == "" || c.AppVersionDate == "" || c.HttpAddress == "" {
		c.InstallTime = time.Now().Unix()
		c.AppVersion = APP_VERSION
		c.AppVersionDate = APP_VERSION

		c.HttpAddress = "0.0.0.0:3030"
		c.DataDirectory = "user/data"
	}

	return c
}

// read config file
func (c *Config) ReadFile() error {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	if len(bytes) > 0 {
		if err = json.Unmarshal(bytes, c); err != nil {
			return err
		}
	}
	return nil
}

// write config file
func (c *Config) WriteFile() error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configFile, bytes, os.ModePerm)
}

// check config file existing
func (c *Config) Exist() bool {
	return com.IsFile(configFile)
}
