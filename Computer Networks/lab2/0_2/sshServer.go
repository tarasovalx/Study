package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	password = flag.String("pwd", "***", "Password")
	host     = flag.String("h", "localhost", "Host")
	port     = flag.String("p", "3000", "Port")
)

func handler(session ssh.Session) {
	var util, res string
	var rawRes []byte
	var args []string
	log.Println("user connected")

	term := terminal.NewTerminal(session, "> ")
	for {
		cmd, err := term.ReadLine()
		log.Println(cmd)
		if err != nil {
			break
		}
		parts := strings.Fields(cmd)
		args = nil
		util = parts[0]
		if len(parts) > 1 {
			args = parts[1:]
		}
		log.Printf("exec %s + %s", util, args)
		if len(parts) > 0 {
			rawRes, err = exec.Command(util, args...).Output()
		} else {
			exec.Command(util)
		}
		res = string(rawRes)
		if res != "" {
			log.Println("res:", res)
			b, err := fmt.Fprintf(session, "%s", util)
			if err != nil {
				log.Println(b, err)
			}
		}
	}
	log.Println("terminal closed")
}

func main() {
	ssh.Handle(handler)
	log.Printf("ssh server started on port %d...", port)
	log.Fatal(ssh.ListenAndServe(
		*host+":"+*port,
		nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			return pass == *password
		}),
	))
}
