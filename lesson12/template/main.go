package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		fmt.Println("listen server failed, err:", err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./login.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed, error:%v", err)
		return
	}
	s1 := &student{
		Name: "Marry",
		Sex:  "female",
		age:  18,
	}
	s2 := make(map[string]string)
	s2["name"] = "Marry"
	s2["sex"] = "female"
	s2["age"] = "18"
	err = t.Execute(w, s1)
	if err != nil {
		fmt.Fprintf(w, "err:%v", err)
		return
	}
}

type student struct {
	Name string
	Sex  string
	age  int
}
