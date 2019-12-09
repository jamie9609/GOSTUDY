package main

import (
	"database/sql"
	"fmt"
	"testing"
)


func TestInit(t *testing.T)  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	_, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		t.Error(err)
	}
}
