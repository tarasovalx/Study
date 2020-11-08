package main

import (
	"flag"
	"log"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

// Точка входа
func main() {

	// флаги для конфигурирования при запуске из терминала
	var (
		root = flag.String("root", "./files", "Root directory to serve")
		user = flag.String("user", "admin", "Username for login")
		pass = flag.String("pass", "123456", "Password for login")
		port = flag.Int("port", 3030, "Port")
		host = flag.String("host", "localhost", "Host")
	)
	flag.Parse()

	factory := &filedriver.FileDriverFactory{
		RootPath: *root,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	//Параметры ftp сервера
	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     *port,
		Hostname: *host,
		Auth:     &server.SimpleAuth{Name: *user, Password: *pass},
	}

	log.Printf("Start ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Server root dir %s", *root)
	log.Printf("Username %v, Password %v", *user, *pass)

	server := server.NewServer(opts)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
