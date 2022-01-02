package type_test

import (
	"testing"
	"fmt"
)

func TestImplicit(t *testing.T)  {
	var a int32 = 10
    var b int64
	type myInt = int64
	var c myInt = 5
	b = int64(a)
    b = c
	t.Log(a,b ,c)
}

func TestPoint(t *testing.T)  {
	var a =  15
	aptr := &a
	//aptr = aptr + 1
	t.Log(a, aptr)
	t.Logf("%T %T", a, aptr)
}

func TestString(t *testing.T)  {
	var s  string 
	for i := range s {
		fmt.Println(i)
	}
	fmt.Println("hello")
	t.Log("*"+s+"*")
	t.Log(len(s))
	if s == "" {
		t.Log("s is nil stirng ")
	}
}

func TestBitClear(t *testing.T)  {
	const (
		Readable = 1 << iota
		Writable
		Excutable
	)

	a := 7 // 0111
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Excutable == Excutable)
	
//异或运算，或者同为1则置0.
	a = a&^Readable
	a = a&^Writable
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Excutable == Excutable)
}