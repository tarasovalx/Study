package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlaffaye/ftp"

	"bytes"
	"log"
)

var connection *ftp.ServerConn
var ftpURL = "students.yss.su:21"

var login = "*****"
var password = "****"
var ftpPath = "./" // path on ftp server


func MakeDir(path string) {
	connection.MakeDir(path)
}

func UploadFile(srcPath string, destPath string) {
	data, _ := ioutil.ReadFile(srcPath)
	reader := bytes.NewBuffer(data)

	if err := connection.Stor(ftpPath+destPath, reader); err != nil {
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

func DownloadFile(srcPath, destPath string) {
	if r, err := connection.Retr(srcPath); err != nil{
		log.Fatal(err)
	} else {
		defer r.Close()
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			panic(err)
		}
		WriteFile(destPath, &buf)
	}
}

func RemoveFile(path string) {
	connection.Delete(path)
}

func Init() {
	connection.MakeDir(ftpPath)
	connection.ChangeDir(ftpPath)
}

func RetriveDir(path string) {
	walker := connection.Walk(path)
	for walker.Next() {
		fmt.Println(walker.Path())
	}
}

func main() {
	var err error
	ftpURLF := flag.String("url", "students.yss.su:21", "Server address")
	loginF := flag.String("user", "username", "Login")
	passwordF := flag.String("pass", "*****", "Password for login")

	flag.Parse()

	ftpURL = *ftpURLF
	login = *loginF
	password = *passwordF

	if connection, err = ftp.Connect(ftpURL); err != nil{
		log.Fatal(err)
	}

	if err = connection.Login(login, password); err != nil {
		log.Fatal(err)
	}

	Init()
	UploadFile("file.txt", "/file.txt")
	DownloadFile("/tat/file.txt", "downloads/file.txt")
	MakeDir("AnotherOne")
	UploadFile("file.txt", "/AnotherOne/file.txt")
	RetriveDir(ftpPath)

	if err := connection.Quit(); err != nil {
		log.Fatal(err)
	}
}
