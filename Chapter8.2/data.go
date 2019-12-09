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

// connect to the Db
func init() {
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

// Get a single post
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
