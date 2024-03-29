package main

import "fmt"

func callerA(c chan string)  {
	c <- "hello world!"
	close(c)
}

func callerB(c chan string)  {
	c <- "hola mundo"
	close(c)
}

func main()  {
	a, b := make(chan string), make(chan string)

	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true

	for ok1 || ok2 {
		select {
		case msg, ok1 = <- a:
			if ok1{
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <- b:
			if ok2{
				fmt.Printf("%s from B\n", msg)
			}
		}
	}

	/*for i := 0; i< 10 ;  i++ {
		time.Sleep(1*time.Microsecond)
		select {
		case msg := <- a:
			fmt.Printf("%s from A\n", msg)
		case msg := <- b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("default")
		}
	}*/
}
