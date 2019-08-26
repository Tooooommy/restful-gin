package ants

import (
	"github.com/panjf2000/ants"
	"restful-gin/config"
)

func InitAnts() {
	cfg := config.Get().Pool
	p, _ := ants.NewPool(cfg.Ants)
	_ = p.Submit(func() {
		println("HELLO")
	})
	p.Running()
}
