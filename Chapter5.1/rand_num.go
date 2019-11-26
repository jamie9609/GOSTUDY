package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func formateDate(t time.Time) string {
	layout := "2019-11-12"
	return t.Format(layout)
}

func process1(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("/Users/didi/go/src/GOSTUDY/Chapter5.1/template/tmpl.html")
	//此处不能是模版的相对路径，需要绝对路径
	rand.Seed(time.Now().Unix())
	//设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	t.Execute(w, rand.Intn(10)>5 )
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello!")
}

func world(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "world!")
}

func week(w http.ResponseWriter, r *http.Request)  {
	funcMap := template.FuncMap{"fdate": formateDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("/Users/didi/go/src/GOSTUDY/Chapter5.1/template/tmpl.html","/Users/didi/go/src/GOSTUDY/Chapter5.1/template/tmpl2.html")
	daysOfWeek1 := []string{"monday","tuesday","wednesday","tuesday","friday","saturday","sunday"}
	//daysOfWeek := []string{}
	t.Execute(w, daysOfWeek1)
}

func process(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("/Users/didi/go/src/GOSTUDY/Chapter5.1/template/tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}

func main1(){
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/process1", process1)
	http.HandleFunc("/process", process)
	//http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/week", week)
	server.ListenAndServe()
}
