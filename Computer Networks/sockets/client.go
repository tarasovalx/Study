package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var PORT = "8080"

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil{
		log.Println("Connect error: ", err.Error())
	}
	for {
		in := bufio.NewReader(os.Stdin)
		fmt.Printf("Write your text:\n")
		msg,_ := in.ReadString('\n')
		fmt.Fprintf(conn, msg + "\n")
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(response)
	}
}