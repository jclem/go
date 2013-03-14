package main

import "fmt"
import "time"

func sleep (channel chan<- string, str string) {
  time.Sleep(time.Second * 2)
  channel <- str
}

func main() {
  c1 := make(chan string)
  go sleep(c1, "result 1")

  select {
  case res := <-c1:
    fmt.Println(res)
  case <-time.After(time.Second * 1):
    fmt.Println("timeout 1")
  }

  c2 := make(chan string)
  go sleep(c2, "result 2")

  select {
  case res := <-c2:
    fmt.Println(res)
  case <-time.After(time.Second * 3):
    fmt.Println("timeout 2")
  }
}
