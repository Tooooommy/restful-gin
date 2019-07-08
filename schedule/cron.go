package schedule

import (
	"CrownDaisy_GOGIN/config"
	"github.com/ivpusic/grpool"
	"time"
)

type Pool struct {
	*grpool.Pool
	tickers []Job
	timers  []Job
}

type Job func()

func NewPool(nws, jql int) *Pool {
	return &Pool{grpool.NewPool(nws, jql), make([]Job, 0), make([]Job, 0)}
}

func (p *Pool) AddTickerJob(cron time.Duration, job Job) {
	p.WaitCount(1)
	go func() {
		defer p.JobDone()
		for range time.NewTicker(cron).C {
			p.JobQueue <- grpool.Job(job)
		}
	}()
}

func (p *Pool) AddTimerJob(cron time.Duration, job Job) {
	p.WaitCount(1)
	go func() {
		defer p.JobDone()
		for range time.NewTimer(cron).C {
			p.JobQueue <- grpool.Job(job)
		}
	}()
}

func (p *Pool) AddTickerImmediateJob(cron time.Duration, job Job) *time.Ticker {
	t := time.NewTicker(cron)
	p.WaitCount(1)
	go func(t *time.Ticker) {
		defer p.JobDone()
		for {
			p.JobQueue <- grpool.Job(job)
			<-t.C
		}
	}(t)
	return t
}

func (p *Pool) AddImmediateJob(job Job) {
	p.WaitCount(1)
	go func() {
		defer p.JobDone()
		p.JobQueue <- grpool.Job(job)
	}()
}

var SchedulePool *Pool

func InitPool() {
	cfg := config.Get().Pool
	SchedulePool = NewPool(cfg.Worker, cfg.Job)
	return
}
