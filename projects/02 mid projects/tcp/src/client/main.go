package client

import (
	"fmt"
	"net"
)


func main() {

	//for connection
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("connetcion error : ", err)
	}

	defer conn.Close()
	_, err = conn.Write([]byte("hello my server"))
	if err != nil {
		fmt.Println("SENDİNG PACKAGE ERROR :", err)
		return
	}
}