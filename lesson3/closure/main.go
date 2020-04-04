package main

import (
	"fmt"
	"strings"
)

func addSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	func1 := addSuffix(".png")
	temp := func1("xiangrikui.png")
	fmt.Println(temp)
}
