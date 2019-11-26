package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User        string
	Threads     []string
}

func writeExample(w http.ResponseWriter, r *http.Request){
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Location","http://www.baidu.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type","application/json")
	post := &Post{
		User:      "jamiezhangming",
		Threads:   []string{"first","second","last"},
	}
	json, _ := json2.Marshal(post)
	w.Write(json)
	fmt.Fprintf(w,"\n")
}

func main1()  {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}