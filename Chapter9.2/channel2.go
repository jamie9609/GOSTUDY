package main

import (
	"fmt"
	"time"
)

func thrower(c chan int)  {
	for i :=0 ; i<5 ; i ++ {
		c <- i
		fmt.Println("threw >>", i)
		//fmt.Printf("threw >> %d", i)
	}
}

func catcher(c chan int)  {
	for i :=0 ; i<5 ; i ++ {
		num := <- c
		fmt.Println("caught <<", num)
		//fmt.Printf("caught << %d", num)
	}
}

func main2()  {
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100*time.Millisecond)
}