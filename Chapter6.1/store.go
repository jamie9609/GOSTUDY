package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post)  {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main1()  {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id:1, Content:"hello world", Author:"jamie"}
	post2 := Post{Id:2, Content:"my name", Author:"jamie2"}
	post3 := Post{Id:3, Content:"is jamiezhangming", Author:"jamie3"}
	post4 := Post{Id:4, Content:"nice to meet you!", Author:"jamie"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["jamie"]{
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["jamie2"]{
		fmt.Println(post)
	}

}