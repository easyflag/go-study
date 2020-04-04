package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:5050")
	if err != nil {
		fmt.Println("dial failed,error:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		n, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("write failed,error:", err)
			return
		}
		fmt.Println(n, "bytes has send")
	}

}
