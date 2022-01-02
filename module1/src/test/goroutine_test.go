package goroutine_test

import (
	"testing"
	"fmt"
	"time"
)

func TestGorouting(t *testing.T)  {
	ch := make(chan string)

    getData(ch)
	go sendData(ch)
	
}

func sendData(ch chan string)  {
	fmt.Printf("in senddata\n")
	ch <- "Hllo"
	
	ch <- "Nobbq"
	ch <- "goxi"
	close(ch)
}

func getData(ch chan string){
	fmt.Printf("before for\n")
	for {
		input , ok := <- ch
		if ok == "" {
			break
		}else{
			fmt.Printf("%s ", input)
		}
	}
}