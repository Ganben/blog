package core

import "github.com/robfig/cron"

// cron struct
type Cron struct {
	*cron.Cron
}

// new cron struct
func NewCron() *Cron {
	return &Cron{cron.New()}
}
