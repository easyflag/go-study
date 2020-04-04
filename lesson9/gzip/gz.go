package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "./file.go.gz"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(fz)
	result := ""
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d has read\n", len(line))
		result += line
	}
	fmt.Println(result)
}
