package helpers

import (
	"fmt"
	"github.com/panjf2000/ants"
	"testing"
)

func TestCoroutine(t *testing.T) {

	p, err := ants.NewPool(100)
	if err != nil {
		panic(err)
	}
	err = p.Submit(func() {
		fmt.Println("Hello")
	})
	if err != nil {
		panic(err)
	}
}
