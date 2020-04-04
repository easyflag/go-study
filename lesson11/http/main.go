package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("dial failed,error:", err)
		return
	}
	defer conn.Close()

	msg := "GET/HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n\r\n"

	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write failed,error:", err)
		return
	}
	fmt.Println(n, "bytes has send")

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed,error:", err)
			break
		}
		fmt.Println("recieve:", string(buf[:n]))
	}
}
