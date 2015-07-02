package model

import (
	"encoding/json"
	"github.com/gofxh/blog/app"
	"github.com/gofxh/blog/app/log"
)

const (
	SETTING_FORMAT_TEXT int8 = 16 // text value
	SETTING_FORMAT_JSON int8 = 26 // json data
)

var Settings map[string]*Setting = make(map[string]*Setting)

// setting struct
type Setting struct {
	Name   string `xorm:"unique(setting-name)"`
	Value  string
	Format int8  `xorm:"not null"`
	UserId int64 `xorm:"unique(setting-name)"`
}

// get string value
func (s *Setting) GetString() string {
	return s.Value
}

// get json value
func (s *Setting) GetJson(value interface{}) error {
	return json.Unmarshal([]byte(s.Value), value)
}

// new default settings
func NewDefaultSetting(uid int64) []*Setting {
	m := make([]*Setting, 0)
	m = append(m, &Setting{"theme", "default", SETTING_FORMAT_TEXT, uid})
	return m
}

// save settings
func SaveSettings(ss []*Setting) error {
	sess := app.Db.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		log.Error("Db|SaveSettings|%s", err.Error())
		return err
	}
	for _, s := range ss {
		sess.Insert(s)
	}
	if err := sess.Commit(); err != nil {
		log.Error("Db|SaveSettings|%s", err.Error())
		return err
	}
	return nil
}

// read settings
func ReadSettingsToGlobal() error {
	settings := make([]*Setting, 0)
	if err := app.Db.Find(&settings); err != nil {
		log.Error("Db|ReadSettingsToGlobal|%s", err.Error())
		return err
	}
	// set to global
	for _, s := range settings {
		Settings[s.Name] = s
	}
	return nil
}
