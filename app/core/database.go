package core

import (
	"github.com/go-xorm/xorm"
	"github.com/gofxh/blog/app/log"
	"github.com/mattn/go-sqlite3"
)

// core database
type Database struct {
	*xorm.Engine
}

// open database file
func NewDatabase(file string) *Database {
	v, _, _ := sqlite3.Version()
	log.Info("Db|Connect|SQLite %s|%s", v, file)
	engine, err := xorm.NewEngine("sqlite3", file)
	if err != nil {
		log.Fatal("Db|Connect|Error|%s", err.Error())
	}
	engine.ShowDebug = true
	engine.ShowInfo = false
	//engine.SetLogger(nil)
	return &Database{engine}
}
