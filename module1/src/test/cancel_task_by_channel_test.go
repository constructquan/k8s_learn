package cancel_task_test

import (
	"testing"
	"fmt"
	"time"
)

func cancel_1(ch chan struct{})  {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}){
	close(ch)
}

func isCancelChannel(cancelch chan struct{}) bool  {
	select {
	case <- cancelch :
		return true
	default:
		return false
	}
}

func TestCancelTask(t *testing.T)  {
	cancelChan := make(chan struct{}, 0)
    for i:=0;i<5;i++{
		go func(i int, cancelCha chan struct{}){
			for{
				if isCancelChannel( cancelCha ){
					break
				}else{
					time.Sleep(time.Millisecond * 5)
				}
			}
			fmt.Printf("%d Canceled\n", i)
		}(i, cancelChan)
	}
	
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}