package core

import (
	"encoding/json"
	"github.com/Unknwon/com"
	"github.com/gofxh/blog/lib/entity"
	"github.com/gofxh/blog/lib/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

// storage
type Storage struct {
	directory string
	lock      sync.Mutex // need locker
}

// new storage with data directory
func NewStorage(dir string) *Storage {
	s := &Storage{
		directory: dir,
	}
	if !com.IsDir(s.directory) {
		os.MkdirAll(s.directory, os.ModePerm)
	}
	return s
}

// save entity
func (s *Storage) Save(e entity.Entity) {
	s.lock.Lock()
	defer s.lock.Unlock()

	bytes, err := json.Marshal(e)
	if err != nil {
		log.Error("Storage|Save|%s|%s", e.SKey(), err.Error())
		return
	}

	file := filepath.Join(s.directory, e.SKey()+".json")
	os.MkdirAll(filepath.Dir(file), os.ModePerm)
	if err = ioutil.WriteFile(file, bytes, os.ModePerm); err != nil {
		log.Error("Storage|Save|%s|%s", e.SKey(), err.Error())
		return
	}
}

// read entity
func (s *Storage) Read(e entity.Entity) {
	s.lock.Lock()
	defer s.lock.Unlock()

	file := filepath.Join(s.directory, e.SKey()+".json")
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("Storage|Read|%s|%s", e.SKey(), err.Error())
		return
	}

	if err = json.Unmarshal(bytes, e); err != nil {
		log.Error("Storage|Read|%s|%s", e.SKey(), err.Error())
		return
	}
}

// check entity existing
func (s *Storage) Exist(e entity.Entity) bool {
	file := filepath.Join(s.directory, e.SKey()+".json")
	return com.IsFile(file)
}

// remove entity
func (s *Storage) Remove(e entity.Entity) {
	s.lock.Lock()
	defer s.lock.Unlock()

	file := filepath.Join(s.directory, e.SKey()+".json")
	if err := os.RemoveAll(file); err != nil {
		log.Error("Storage|Remove|%s|%s", e.SKey(), err.Error())
	}
}

// walk saving and same entity data
func (s *Storage) Walk(e entity.Entity, fn func(interface{})) {
	// get directory
	file := filepath.Join(s.directory, e.SKey()+".json")
	dir := filepath.Dir(file)

	// walk files
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if ext := filepath.Ext(path); ext != ".json" {
			return nil
		}
		nv := reflect.New(reflect.TypeOf(e).Elem()).Interface()
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(bytes, nv); err != nil {
			return err
		}
		if fn != nil {
			fn(nv)
		}
		return nil
	}
	filepath.Walk(dir, walkFn)
}
