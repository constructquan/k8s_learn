package  main

import (
	"fmt"
	"time"

)

func ProdcerData(intChan chan int){
	ticker := time.NewTicker(1 * time.Second)
	num := 1
	go func(iChan chan int){
		for  {
			select {
			case <- ticker.C:
				iChan <- num
				num++
			default:
				continue
			}
			
		}
		
	}(intChan)
}

func ConsumerData(intChan chan int){
	ticker := time.NewTicker(time.Second * 1)
	go func (iChan chan int){
		for {
			select {
			case <- ticker.C:
				fmt.Printf("取出数据： %d\n", <-iChan)
			}
		}
		
	}(intChan)
}

func main(){
	fmt.Print("Start\n")
	
	intChan := make(chan int, 10)
	

	go func (intCh chan int)  {
		ticker := time.NewTicker(1 * time.Second)
		//defer ticker.Stop()
		for i:=1;i<12;i++{
			<- ticker.C
			intCh<-i
		}
	}(intChan)

	go func (intCh chan int)  {
		ticker := time.NewTicker(1 * time.Second)
		for{
			select{
			case <-ticker.C:
				fmt.Printf("数据为：%d\n", <-intCh)
			default:
				continue
			}
		}
	}(intChan)

	for{
		
	}
}