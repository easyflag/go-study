package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("run", (i + 1), "s")
	}
}

func task2() {
	for i := 20; i < 30; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("run", (i + 1), "s")
	}
}

func main() {
	go task1()
	go task2()
	time.Sleep(time.Second * 15)
	fmt.Println("over")
}
