package interface_test

import "testing"

type Programer interface {
	WriteHelloWorld()string
}	

type GoProgramer struct {
}

//接口实现的方式之一：
func (p *GoProgramer) WriteHelloWorld()string  {
	return "Hello world from GoProgramer"
}

func TestClient(t *testing.T)  {
	var p Programer
	p = new(GoProgramer)
	t.Log(p.WriteHelloWorld())
}