package main

///Импорт библиотек
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

///Конфигурация аргументов командной строки по умолчанию
var (
	user     = flag.String("u", "tarasov", "User name")
	password = flag.String("pwd", "1234567890", "Password")
	host     = flag.String("h", "lab2.posevin.com", "Host")
	port     = flag.String("p", "22", "Port")
)

var client *ssh.Client

///Точка входа в программу
func main() {
	flag.Parse()

	///Конфигурация ssh клиента
	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.Password(*password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var err error

	client, err = ssh.Dial("tcp", *host+":"+*port, config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to %s:%s ...", *host, *port)

	reader := bufio.NewReader(os.Stdin)

	///Считываем ввод до ctrl+C или до exit
	for {
		fmt.Printf("> ")
		cmd, _ := reader.ReadString('\n')
		execcmd(strings.TrimSpace(cmd))
		if cmd == "exit" {
			return
		}
	}
}

func execcmd(cmd string) {
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
