package main

import (
	"fmt"
	"time"
)

func showTime() {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}

func taskTimer() {
	start := time.Now().Nanosecond()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().Nanosecond()
	value := (end - start) / 1000
	fmt.Printf("任务耗时%dms", value)
}

func main() {
	showTime()
	taskTimer()
}
