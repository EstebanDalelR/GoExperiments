package main

import (
"fmt"
"time"
)

func main()  {
  chan1 := make(chan string)
  chan2 := make(chan string)

  go func ()  {
    for {
      chan1 <- "From1"
      time.Sleep(time.Second * 2)
    }
  }()
  go func ()  {
    for {
      chan2 <- "From2"
      time.Sleep(time.Second * 3)
    }
  }()
  go func ()  {
    for {
      select {
      case msg1 := <- chan1:
        fmt.Println(msg1)
      case msg2 := <- chan2:
        fmt.Println(msg2)
      case <- time.After(time.Second):
        fmt.Println("Timeout")
      }
    }
  }()

  var inp string
  fmt.Scanln(&inp)
}
