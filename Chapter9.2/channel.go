package main

import (
	"fmt"
	"time"
)

func printNumber2(w chan bool)  {
	for i := 0; i<10 ;i++  {
		time.Sleep(1*time.Microsecond)
		fmt.Printf("%d", i)
	}
	w <- true
}

func printLetters2(w chan bool)  {
	for i := 'A'; i<10 + 'A';i++  {
		time.Sleep(1*time.Microsecond)
		fmt.Printf("%c", i)
	}
	w <- true
}

func main1()  {
	w1 := make(chan bool)
	w2 := make(chan bool)
	go printNumber2(w1)
	go printLetters2(w2)
	<- w1
	<- w2
}
