package channel_first_test

import (
	"testing"
)

func TestChannelFirst(t *testing.T)  {
	var iChan chan int = make(chan int, 3)
	t.Logf("iChan is %T", iChan) 
	t.Logf("iChan is %v", iChan)

	iChan <- 10
	num := 100
	iChan <- num

	t.Logf("iChan len: %d, cap: %d", len(iChan), cap(iChan))

	//取数据
	num2 := <- iChan
	t.Logf("取出数据： %d", num2)
	//查看长度，容量是否变化
	t.Logf("iChan len: %d, cap: %d", len(iChan), cap(iChan))
	iChan<- num2
 	//关闭管道   
	close(iChan)
	for {
		v, ok := <- iChan
		if !ok{
			break
		}
		t.Logf("取出数据来： %d", v)
	}
	
}
	