package mutex_test

import (
	"sync"
	"testing"
	"fmt"
)

func Lock()  {
	lock := sync.Mutex 
	for i:=1;i <4; i++{
		lock.Lock()
		defer lock.Unlock()
		fmt.Printf("Lock: %d\n", i)
	}
}

func TestMutex(t *testing.T)  {
	go Lock()
}


