package producer_test

import (
	"time"
	"testing"
	"fmt"
	"errors"
)

func dataProducer(iCh chan int, n int)  error  {	
	select{
		case iCh <- n:
			fmt.Printf("插入数据：%d\n", n)
			return nil
		case <-time.After(time.Second * 1):
			return errors.New("time out\n")
		default:
			return errors.New("channel is  full\n")
	}
}

func dataConsumer(iCh chan int)  error {
	select{
		case ret:= <- iCh:
			fmt.Printf("取出数据：%d\n", ret)
			return  nil
		case <-time.After(time.Second * 1):
			return errors.New("time out\n")
		default:
			return errors.New("channel is empty\n")
	}
}

func TestProducerConsumer(t *testing.T)  {
	intChan := make(chan int, 10)
	ticker := time.NewTicker( time.Second * 1)
	num:=12
	
    for i:=0;i<num;i++{
		<-ticker.C
		if err:=dataProducer(intChan, i);err !=nil{
			fmt.Print(err.Error())
		}
	}

	for {
		<-ticker.C
		if err:= dataConsumer(intChan);err !=nil{
			fmt.Print(err.Error())
		}
	}

}