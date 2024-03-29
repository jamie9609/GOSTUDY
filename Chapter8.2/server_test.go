package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, err := http.NewRequest("GET","/post/1",nil)
	mux.ServeHTTP(writer, request)

	if err != nil{
		t.Error(err)
	}

	if writer.Code != 200{
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post

	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1{
		t.Errorf("cannot retrieve json post")
	}
}
