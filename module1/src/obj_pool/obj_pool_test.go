package main

import (
	"testing"
	"fmt"
	"time"
)

func TestObjPool(t *testing.T)  {
	pool := NewObjPool(10)
	for i:=0; i<11;i++{
		if v, err := pool.GetObj(time.Second * 1); err != nil{
			t.Error(err)
		}else{
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil{
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}