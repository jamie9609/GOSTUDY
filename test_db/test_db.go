package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "192.168.43.184"
	port     = 3307
	user     = "root"
	password = "z19930708m"
	dbname   = "chitchat"
)


func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer Db.Close()

	//验证连接
	if err := Db.Ping(); err != nil{
		fmt.Println("Opon database fail")
		return
	}
	/*	err = Db.Ping()
		if err != nil {
			panic(err)
		}*/
	fmt.Println("Successfully connected!")


}