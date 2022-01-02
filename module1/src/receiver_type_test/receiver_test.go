package receiver_test

import (
	"fmt"
	"testing"
	"reflect"
)


func dumpMethodSet(i interface{}) {
    dynTyp := reflect.TypeOf(i)

    if dynTyp == nil {
        fmt.Printf("there is no dynamic type\n")
        return
    }

    n := dynTyp.NumMethod()
    if n == 0 {
        fmt.Printf("%s's method set is empty!\n", dynTyp)
        return
    }

    fmt.Printf("%s's method set:\n", dynTyp)
    for j := 0; j < n; j++ {
        fmt.Println("-", dynTyp.Method(j).Name)
    }
    fmt.Printf("\n")
}

type T struct{

}

type Interface interface{
	M1()
	M2()
}

func ( T)M1()  {
	
}

func ( T )M2(){

}


type S T

func TestReceiver(t *testing.T)  {
	var  ty T 
	var pty *T
	dumpMethodSet(ty)
	dumpMethodSet(pty )

	var st S
	var pst *S
	dumpMethodSet(st)
	dumpMethodSet(pst)

}