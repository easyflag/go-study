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
	if r.Method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed, error:%v", err)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("username:%s\n", username)
		fmt.Printf("password:%s\n", password)
		if username == "953284879" {
			if password == "sb1063367806" {
				fmt.Fprintf(w, "user %s login success", username)
			} else {
				fmt.Fprintf(w, "user %s password error!", username)
			}
		}
	}
}
