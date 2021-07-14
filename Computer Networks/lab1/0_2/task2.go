package main

import (
	"flag"
	"log"

	fileDriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

func main() {
	var (
		root = flag.String("root", "./files", "Root directory to serve")
		user = flag.String("user", "admin", "Username for login")
		pass = flag.String("pass", "123456", "Password for login")
		port = flag.Int("port", 3030, "Port")
		host = flag.String("host", "localhost", "Host")
	)
	flag.Parse()

	factory := &fileDriver.FileDriverFactory{
		RootPath: *root,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     *port,
		Hostname: *host,
		Auth:     &server.SimpleAuth{Name: *user, Password: *pass},
	}

	log.Printf("Start ftp ftpServer on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Server root dir %s", *root)
	log.Printf("Username %v, Password %v", *user, *pass)

	ftpServer := server.NewServer(opts)
	if err := ftpServer.ListenAndServe(); err != nil {
		log.Fatal("Error starting ftpServer:", err)
	}
}
