package main

import (
	"fmt"
	"net"
)

func main() {
	addr1 := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 2020,
	}
	conn, err := net.ListenUDP("udp", addr1)
	if err != nil {
		fmt.Println("listen failed,error:", err)
	}
	defer conn.Close()

	response(conn)
}

func response(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Read failed,error:", err)
			continue
		}
		fmt.Printf("recieve from %v:%s", addr, buf[:n])

		n, err = conn.WriteToUDP([]byte("hello client"), addr)
		if err != nil {
			fmt.Println("Write failed,error:", err)
			continue
		}
		fmt.Println(n, "bytes has send")
	}
}
