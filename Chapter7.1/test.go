package Chapter7_1

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	High  bool
	Sex   string
	Class *Class `json:"class"`
}


type StudentRead struct {
	Name  string `json:"name"`
	Age   int
	High  bool
	Sex   string
	ClassRead *Class `json:"ClassRead"`
}

type ClassRead struct {
	Name  string
	Grade int
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	stu := Student{
		"zhangsan",
		18,
		true,
		"男",
		&Class{},
	}

	cla := new(Class)
	cla.Name = "一班"
	cla.Grade = 3
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil{
		fmt.Println("生成json错误")
	}
	fmt.Println(string(jsonStu))

	data:="{\"name\":\"张三\",\"Age\":18,\"high\":true,\"sex\":\"男\",\"ClassRead\":{\"Name\":\"1班\",\"Grade\":3}}"

	str:=[]byte(data)
	sturead:= StudentRead{}
	err2 := json.Unmarshal(str, &sturead)
	if err2 != nil{
		fmt.Printf("生成json错误,类型为%s", err2)
	}
	fmt.Println(sturead)

}