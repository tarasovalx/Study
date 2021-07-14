package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

var (
	user     = flag.String("u", "login", "User name")
	password = flag.String("pwd", "pass", "Password")
	host     = flag.String("h", "host", "Host")
	port     = flag.String("p", "22", "Port")
)

var client *ssh.Client

func main() {
	var err error
	flag.Parse()
	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}


	if client, err = ssh.Dial("tcp", *host+":"+*port, config);err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to %s:%s ...", *host, *port)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		cmd, _ := reader.ReadString('\n')
		executeCommand(strings.TrimSpace(cmd))
		if cmd == "exit" {
			os.Exit(0)
		}
	}
}

func executeCommand(cmd string) {
	session, _ := client.NewSession()
	resp, err := session.Output(cmd)
	if err != nil {
		log.Println(err)
	}
	if string(resp) != "" {
		log.Println(string(resp))
	}
	defer session.Close()
}
