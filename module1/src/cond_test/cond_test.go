package cond_test

import (
	"testing"
	"fmt"
	"sync"
)

type Item = int

type Queue struct{
	items []Item
	itemAdded sync.Cond
}

func NewQueue() *Queue{
	q := new(Queue)
	q.itemAdded.L  = &sync.Mutex{}  //绑定 cond 的锁

	return q
}

func ( q *Queue) PutItem(item Item){
	q.itemAdded.L.Lock()
	defer q.itemAdded.L.Unlock()

	q.items = append(q.items, item)
	q.itemAdded.Signal() //当Queue加入数据成功后，就调用Signal() 通知
}

func (q *Queue) GetItem(num int) []Item  {
	q.itemAdded.L.Lock()
	defer q.itemAdded.L.Unlock()
	for len(q.items)<num {
		q.itemAdded.Wait()
	}

	items := q.items[:num:num]
	q.items = q.items[num:]
	
	return items
}

func TestCond(t *testing.T)  {
	q := NewQueue()
	var wg sync.WaitGroup
	
	//起 10 个协程，分别去取队列里的数据
	for num:=10;num>0;num--{
		wg.Add(1)
		go func(n int){
			items:=q.GetItem(n)
			fmt.Printf("%d 取出队列数据： %d\n",n, items)
			wg.Done()
		}(num)
	}

	
	//放入100个数
	for n:=0;n<100;n++{
		q.PutItem(n)
	}
	wg.Wait()
}