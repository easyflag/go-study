package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//file, err := os.Open(`C:/Users/24213/go project/src/github.com/
	//go-study/lesson8/calculater/logic.go`)
	file, err := os.Open("./bufio.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content := ""
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d bytes has read\n", len(line))
		content += line
	}
	fmt.Println(content)
}
