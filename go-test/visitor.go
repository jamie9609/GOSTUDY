package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = map[string]int{}

func handleHello(w http.ResponseWriter, r *http.Request )  {
	name := r.FormValue("name")
	counter[name]++

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") + "'>Welcome!</h1> <p>Name: " + name + "</p> <p>Count: " + fmt.Sprint(counter[name]) + "</p>"))
}

func main()  {
	http.HandleFunc("/hello", handleHello)
	log.Fatal(http.ListenAndServe(":8080",nil))
}