package main

import (
	"fmt"
	"time"
)

func printNumbers1()  {
	for i := 0; i< 10; i++{
		time.Sleep(1*time.Microsecond)
		fmt.Printf("%d", i)
	}
	fmt.Print("\n")
}

func printLetters1()  {
	for i := 'A'; i< 10+'A'; i++ {
		time.Sleep(1*time.Microsecond)
		fmt.Printf("%c", i)
	}
	fmt.Print("\n")
}

func print1(){
	printNumbers1()
	printLetters1()
}


func goprint1()  {
	go printNumbers1()
	go printLetters1()
}

func main()  {
	print1()
	goprint1()

}