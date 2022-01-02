package channel_rw_test

import (
	"testing"
	"fmt"
	"time"
)
func ReadData(iChan chan int){
	for {
		v,ok := <- iChan
		if ok {
			fmt.Printf("管道取出的数据： %d\n", v)
		}else{
			break
		}
		
	}

}


func WriteData(iChan chan int)  {
	for i:=1; i<20; i++{
		iChan <- i
	}
	close(iChan)
}
func TestChannelReadWrite(t *testing.T)  {
	iChan := make(chan int)
	
	go WriteData(iChan)
	go ReadData(iChan)
	
	t.Log("Done")
	time.Sleep(time.Second * 1)

}