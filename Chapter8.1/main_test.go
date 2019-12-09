package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T)  {
	post, err := decode("post.json")
	if err != nil{
		t.Error(err)
	}
	if post.Id != 1{
		t.Error("wrong id, was except 1 but got", post.Id)
	}
	if post.Content != "Hello World!"{
		t.Error("Post content is not the same as post.json", post.Content)
	}
}

func TestEncode(t *testing.T)  {
	t.Skip("skipping")
}

func TestLongRunningTest(t *testing.T){
	if testing.Short(){
		t.Skip("skipping long-running test in short mode")
	}
	time.Sleep(10*time.Second)
}