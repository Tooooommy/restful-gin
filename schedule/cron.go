package schedule

import "github.com/roylee0704/gron"

var Cron *gron.Cron

func init() {
	Cron = gron.New()
}
