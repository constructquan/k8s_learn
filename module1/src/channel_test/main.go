package main

import (
   "fmt"
   "time"
)

func main(){
   messages := make(chan int, 10)
   done := make(chan bool)
   secondDone := make(chan bool)
   defer close(messages)
   go func() {
      for {
         select {
            case <-done:
               fmt.Println("chan receive close .....")
               return
            default:
               fmt.Printf("receive message :%d\n", <-messages)
         }
      }

   }()

   ticker := time.NewTicker(1 * time.Second)
   i := 0
   for _ = range ticker.C {
      select {
         case <-secondDone:
            fmt.Println("chan send close .....")
            return
         default:
            i = i + 1
            messages <- i
      }
   }
   time.Sleep(5 * time.Second)
   close(secondDone)
   close(done)

   time.Sleep(1 * time.Second)
   fmt.Println("main process end.....")
   
}