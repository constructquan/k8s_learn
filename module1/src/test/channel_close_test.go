package channel_close_test

import (
//	"time"
	"testing"
	"sync"
	"fmt"
)

func DataWrite(iChan chan int, wg *sync.WaitGroup)  {
	go func (){
		for i:=0;i<15;i++ {
			iChan <- i
		}
		close(iChan)
		//往关闭的channel发送消息，会报错
		//iChan <- 3 
		wg.Done()
	}()
	
}

func DataRead(iChan chan int, wg *sync.WaitGroup)  {
	go func(){
		for {
			if v,ok := <- iChan;ok {
				fmt.Printf("读取数据：%d\n", v)
			}else{
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T)  {
	var wg sync.WaitGroup

	var iChan chan int
	iChan = make(chan int, 1)

	wg.Add(1)
	DataWrite(iChan, &wg)
	wg.Add(1)
	DataRead(iChan, &wg)
	wg.Add(1)
	DataRead(iChan, &wg)

	wg.Wait()
	t.Log("Done in main thread")

}