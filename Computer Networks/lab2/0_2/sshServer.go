package main

///Импорт библиотек
import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

///Конфигурация аргументов командной строки по умолчанию
var (
	user     = flag.String("u", "tarasov", "User name")
	password = flag.String("pwd", "admin", "Password")
	host     = flag.String("h", "localhost", "Host")
	port     = flag.String("p", "9034", "Port")
)

/// Обработчик запросов
func handler(session ssh.Session) {
	var util, res string
	var rawres []byte
	var args []string
	log.Println("user connected")
	///Создаем терминал
	term := terminal.NewTerminal(session, "> ")

	///Считаем ввод до ctrl+C или exit
	for {
		cmd, err := term.ReadLine()
		log.Println(string(cmd))
		if err != nil {
			break
		}
		parts := strings.Fields(string(cmd))
		args = nil
		util = parts[0]
		if len(parts) > 1 {
			args = parts[1:]
		}
		log.Printf("exec %s + %s", util, args)
		if len(parts) > 0 {
			rawres, err = exec.Command(util, args...).Output()
		} else {
			exec.Command(util)
		}
		res = string(rawres)
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

///Точка входа в программу
func main() {
	ssh.Handle(handler)

	log.Println("ssh server started on port 9034...")
	///Запус сервера, логгирование в случае ошибки
	log.Fatal(ssh.ListenAndServe(
		*host+":"+*port,
		nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			return pass == *password
		}),
	))
}
