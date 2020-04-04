package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

func main() {
	students := make([]*student, 0, 10)
	for i := 0; i < 10; i++ {
		stu := &student{
			Person: Person{
				Name: fmt.Sprintf("stu%d", rand.Intn(100)),
				Age:  uint8(18 + rand.Intn(3)),
				Sex:  "male",
			},
			Class: "1554",
			Score: uint8(30 + rand.Intn(50)),
		}
		students = append(students, stu)
	}

	//write(students)
	read(students)
}

func write(students []*student) {
	data, err := json.Marshal(students)
	if err != nil {
		fmt.Printf("marshal failed,error:%v", err)
	}
	err = ioutil.WriteFile("./json.log", data, 0755)
	if err != nil {
		fmt.Printf("write failed,error:%v", err)
	}
}

func read(students []*student) {
	data, err := ioutil.ReadFile("./json.log")
	if err != nil {
		fmt.Printf("read failed,error:%v", err)
	}
	err = json.Unmarshal(data, &students)
	if err != nil {
		fmt.Printf("Unmarshal failed,error:%v", err)
	}

	for _, v := range students {
		fmt.Printf("%#v\n", v)
	}
}

type student struct {
	Person
	Class string
	Score uint8
}

type Person struct {
	Name string
	Age  uint8
	Sex  string
}
