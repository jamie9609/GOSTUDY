package Chapter3_1

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello!")
}

func world(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "world!")
}

func main1(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}