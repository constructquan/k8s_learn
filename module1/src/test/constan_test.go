package constant_test

import (
	"testing"
)

const (
	Monday = 1 + iota //如果不使用 iota，那么下面省略赋值的，就仍然默认是 1.
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

//用二进制位来表示可读、可写、可执行
const (
	Readable = 1 << iota
	Writable
	Excutable
)

func TestConstantTry(t *testing.T)  {
	t.Log("星期天数测试：")
	t.Log(Monday, Tuesday, Wednesday)

	
}

func TestConstantBit(t *testing.T){
	t.Log("读写执行测试：（按位表示）")
	t.Log(Readable, Writable, Excutable)

//	a := 7 //0111
	a := 2 //0010  可写
	
	t.Log(a&Writable == Writable)
	t.Log(a&Excutable == Excutable)
	t.Log(a&Readable == Readable)
}