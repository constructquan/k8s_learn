package struct_emb_test

import (
	"fmt"
	"testing"
	"strings"
	"io"
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

type I interface{
	M1()
	M2()
}

type MyInt int

func (n *MyInt) Add(m int) {
    *n = *n + MyInt(m)
}

type rt struct {
    a int
    b int
}

type S struct {
    *MyInt
    rt
    io.Reader
    s string
    n int
	I
}

func( S ) M3(){}

func TestStructEm(t *testing.T) {
    m := MyInt(17)
    r := strings.NewReader("hello, go")
    s := S{
        MyInt: &m,
        rt: rt{
            a: 1,
            b: 2,
        },
        Reader: r,
        s:      "demo",
    }

    var sl = make([]byte, len("hello, go"))
    s.Read(sl)
    fmt.Println(string(sl)) // hello, go
    s.Add(5)
    fmt.Println(*(s.MyInt)) // 22

	var ps *S
	dumpMethodSet(s)
	dumpMethodSet(ps)
}