package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var PORT = "8080"

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil{
		log.Fatal("Unable to listen", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Unable to open port", err.Error())
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Write([]byte(strings.ToUpper(msg) + "\n"))
	}
}