package main

import (
	"fmt"
	"golang-hacking/03_TCP_Scanner/mynmap"
	"net"
	"sync"
)

func ScanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	IP := "scanme.nmap.org"
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	fmt.Printf("[+] Connection established... PORT %v %v\n", port, connection.RemoteAddr().String())
	connection.Close()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go ScanPort(i, &wg)
	}
	wg.Wait()
	mynmap.Nmap()
}
