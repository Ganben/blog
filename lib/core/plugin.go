package core

type Plugin interface {
	Init()           // init
	Enable()         // enable
	Disable()        // disable
	IsEnabled() bool // enabled status
	Name() string    // name
}

type Plugins map[string]Plugin

func NewPlugins() *Plugins {
	m := Plugins(make(map[string]Plugin))
	return &m
}

func (p *Plugins) Register(pl Plugin) {
	(*p)[pl.Name()] = pl
	pl.Init()
}
