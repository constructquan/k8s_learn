package first_response

import (
	"testing"
	"time"
	"fmt"
	"runtime"
)

func RunTask(id int)string  {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("从任务 %d 返回了结果。", id)
}

func FirstResponse() string {
	num := 10
	var ch chan string
	ch = make(chan string, num)
	
	for i:=0; i<num; i++{
		go func(i int){
			ch <- RunTask(i)
		}(i)
	}
	//如果是没有缓冲的channel，当channel满时，会阻塞。当channel没有线程读取时，也会阻塞
	return <- ch
}

func AllResponse() string {
	num := 10
	var ch chan string
	ch = make(chan string, num)
	
	for i:=0; i<num; i++{
		go func(i int){
			ch <- RunTask(i)
		}(i)
	}
	// return <- ch 

	finallRet := ""
	for j:=0;j<num;j++{
		finallRet += <- ch +"\n"
	}
	return finallRet
}

func TestFirstResponse(t *testing.T)  {
	t.Logf("Before--协程数：%d\n", runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Logf("After--协程数： %d\n", runtime.NumGoroutine())
	time.Sleep(time.Second * 1)
}