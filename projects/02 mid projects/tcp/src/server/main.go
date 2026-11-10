package server

import (
	"fmt"
	"net"
)

func main() {
	listener, err :=  net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listening error :", err)
	}
	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accepting error :", err)
			continue
		}
		go handleConnection(conn)
	}
	
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("CONNETİNG NEW CLİENT:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("ERROR of reading :", err)
			return
		}
		
		fmt.Printf("recived packages : %s\n", string(buffer[:n]))
	}
}