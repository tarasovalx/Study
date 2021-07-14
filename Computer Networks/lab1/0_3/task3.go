package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlaffaye/ftp"

	"flag"
	"log"

	fileDriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

var login, password = "***", "***"

func ConnectAndAuth(url, login, password string) *ftp.ServerConn {
	connection, err := ftp.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Login(login, password)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func RunServer() *server.Server {
	flag.Parse()

	if *root == "" {
		log.Fatalf("Please set a root to serve with -root")
	}

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

	log.Printf("Starting ftp ftpServer on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Server root dir %s", *root)
	log.Printf("Username %v, Password %v", *user, *pass)
	ftpServer := server.NewServer(opts)
	if err := ftpServer.ListenAndServe(); err != nil {
		log.Fatal("Error starting ftpServer:", err)
	}
	return ftpServer
}

var (
	root = flag.String("root", "./files", "Root directory to serve")
	user = flag.String("user", "admin", "Username for login")
	pass = flag.String("pass", "123456", "Password for login")
	port = flag.Int("port", 3030, "Port")
	host = flag.String("host", "localhost", "Host")
)

func MakeDir(path string, connection *ftp.ServerConn) {
	connection.MakeDir(path)
}

func UploadFile(srcPath string, destPath string, connection *ftp.ServerConn) {
	data, err := ioutil.ReadFile(srcPath)

	fmt.Println(data)
	reader := bytes.NewBuffer(data)

	err = connection.Stor(destPath, reader)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFile(path string, data *[]byte) {
	file, _ := os.Create(path)
	_, err := file.Write(*data)
	if err != nil {
		log.Fatal(err)
	}
}

func DownloadFile(srcPath, destPath string, connection *ftp.ServerConn) {
	r, err := connection.Retr(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	WriteFile(destPath, &buf)
}

func RemoveFile(path string, connection *ftp.ServerConn) {
	connection.Delete(path)
}

func Init(connection *ftp.ServerConn, rootPath string) {
	connection.MakeDir(rootPath)
	connection.ChangeDir(rootPath)
}

func RetriveDir(path string, connection *ftp.ServerConn) {
	walker := connection.Walk(path)
	for walker.Next() {
		fmt.Println(walker.Path())
	}
}

func main() {
	_ = RunServer()
	localFtp := ConnectAndAuth("localhost:3030", login, password)
	remoteFtp := ConnectAndAuth("students.yss.su:21", "ftpuser", "*****")
	Init(localFtp, "")
	Init(remoteFtp, "./tat")
}
