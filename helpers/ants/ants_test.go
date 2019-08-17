package ants

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"testing"
)

func TestCoroutine(t *testing.T) {
	num := 10000000
	defer ants.Release()
	wg := sync.WaitGroup{}
	wg.Add(num)
	p, err := ants.NewTimingPoolWithFunc(100, 100, func(i interface{}) {
		defer wg.Done()
		fmt.Println(i)
	})
	defer p.Release()
	if err != nil {
		t.Error(t)
	}
	for i := 0; i < num; i++ {
		_ = p.Invoke(i)
	}
	wg.Wait()
}
