package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func checkPort(host string, port int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	target := fmt.Sprintf("%s:%d", host, port)
	 //debug time out
	conn, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		return
	} 
	conn.Close()

	results <- port

}

func main() {
	fmt.Printf("Welcome Port Checker! \n")
	host := "127.0.0.1"
	startPort := 1
	endPort := 1024
	fmt.Printf("%s port scanning(%d-%d)....\n", host, startPort, endPort)
	var wg sync.WaitGroup
	results := make(chan int, endPort-startPort+1)
	

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go checkPort(host, port, &wg, results)
	}
	wg.Wait()
	close(results)

	fmt.Println("\nopen ports:")
	count := 0
	for port := range results {
		fmt.Printf("port %d is open\n", port)
		count++
	}
	if count == 0 {
		fmt.Println("all ports is closed")
	}
}
