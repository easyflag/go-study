package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/index", login)
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		fmt.Println("listen server failed, err:", err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintf(w, "load index.html failed, error:%v", err)
		return
	}

	students := make([]*student, 0, 30)
	for i := 0; i < 30; i++ {
		stu := &student{
			Name: fmt.Sprintf("Mike%d", rand.Intn(10000)),
			Sex:  "male",
			Age:  rand.Intn(100),
		}
		students = append(students, stu)
	}

	err = t.Execute(w, students)
	if err != nil {
		fmt.Fprintf(w, "err:%v", err)
		return
	}
}

type student struct {
	Name string
	Sex  string
	Age  int
}
