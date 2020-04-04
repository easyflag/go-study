package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start listen-----------")
	listen, err := net.Listen("tcp", "0.0.0.0:5050")
	if err != nil {
		fmt.Println("listen failed,error:", err)
		return
	}

	for {
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,error:", err)
			continue
		}
		go process(connect)
	}

}

func process(connect net.Conn) {
	var buf [1024]byte
	defer connect.Close()
	for {
		n, err := connect.Read(buf[:])
		if err != nil {
			fmt.Println("read failed,error:", err)
			break
		}
		fmt.Println("recieved:", string(buf[:n]))
	}
}
