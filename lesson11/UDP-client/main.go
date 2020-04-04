package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	laddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3030,
	}
	raddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 2020,
	}
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		fmt.Println("Dial failed,error:", err)
	}
	defer conn.Close()

	hello(conn)
}

func hello(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Write([]byte("hello server"))
		if err != nil {
			fmt.Println("Write failed,error:", err)
			continue
		}
		fmt.Println(n, "bytes has send")

		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Read failed,error:", err)
			continue
		}
		fmt.Printf("recieve from %v:%s\n", addr, buf[:n])

		time.Sleep(time.Second * 10)
	}
}
