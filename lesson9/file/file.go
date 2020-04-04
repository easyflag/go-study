package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//file, err := os.Open(`C:/Users/24213/go project/src/github.com/
	//go-study/lesson8/calculater/logic.go`)
	file, err := os.Open("./flie.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := make([]byte, 128, 128)
	//var buf [128]byte
	content := ""
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d bytes has read\n", n)
		content += string(buf[:n])
	}
	fmt.Println(content)
}
