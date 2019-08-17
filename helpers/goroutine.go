package helpers

import (
	"CrownDaisy_GOGIN/config"
	"github.com/panjf2000/ants"
)

func InitAnts() {
	cfg := config.Get().Pool
	p, _ := ants.NewPool(cfg.Ants)
	_ = p.Submit(func() {
		println("HELLO")
	})
}
