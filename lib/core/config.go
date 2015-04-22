package core

import (
	"cloud-server/lib/core/log"
	"encoding/json"
	"github.com/Unknwon/com"
	"io/ioutil"
	"os"
	"time"
)

var configFile string = "config.json"

type Config struct {
	InstallTime    int64  `json:"install_time"`
	AppVersion     string `json:"app_version"`
	AppVersionDate string `json:"app_version_date"`

	HttpAddress   string `json:"http_address"`
	DataDirectory string `json:"data_directory"`
}

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
		c.DataDirectory = "dat"
	}

	return c
}

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

func (c *Config) WriteFile() error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configFile, bytes, os.ModePerm)
}
