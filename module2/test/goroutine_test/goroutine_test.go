package goroutine_test

import (
	"fmt"
	"runtime"
	"testing"
	// "time"
)

func TestGoroutine(t *testing.T)  {
	name := "Harry"
	go func() {
		fmt.Printf("Hello %s.\n", name)
	}()
	name = "Boom"
	runtime.Gosched()
	// time.Sleep(time.Second * 1)
	

}