package schedule

import (
	"fmt"
	"github.com/roylee0704/gron"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	ch := make(chan string, 0)
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Second), func() {
		fmt.Println("Hello")
	})
	c.AddFunc(gron.Every(2*time.Second), func() {
		fmt.Println("world")
	})
	c.Start()
	<-ch
}
