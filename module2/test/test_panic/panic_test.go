package test_panic

import (
	"fmt"
	"testing"
)

func foo(){
	println("call foo")
	bar()
	println("exit foo")
}

func bar(){
	defer func(){
		if e := recover(); e != nil{
			fmt.Println("recover the panic: ", e)
		}
	}()

	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")

}

func zoo(){
	println("call zoo")
	println("exit zoo")
}

func TestPanicSchedual(t *testing.T){
	println("call main")
	foo()
	println("exit main")
}
