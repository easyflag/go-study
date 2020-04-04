package main

import (
	"fmt"
)

func v1() {

}

func main() {
	var str []string = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		str = append(str, fmt.Sprintf("o%d", i))
		fmt.Printf("addr:%p  len:%d  cap:%d\n", str, len(str), cap(str))
	}

	for i := range str {
		fmt.Printf("%s,", str[i])
	}
}
