package singleton_test

import (
	"testing"
	"fmt"
	"sync"
	"unsafe"
)

type Singleton struct{
	data string
}

var once sync.Once
var singleInstance *Singleton

func GetSingletonObj() *Singleton{
	once.Do(func ()  {
		fmt.Println("create singleton obj.")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingleton(t *testing.T)  {
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func ()  {
			obj := GetSingletonObj()
			fmt.Printf("%x \n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
	
}