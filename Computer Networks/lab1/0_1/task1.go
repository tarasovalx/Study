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
var timeOut = 5

var login = "ftpiu8"
var password = "3Ru7yOTA"
var ftpPath = "/tat"
var localPath = "/downloads"

// Создает папку на удаленном сервере
func MakeDir(path string) {
	connection.MakeDir(path)
}

// Загрузка файла на сервер. Принимает локальный путь,
// и путь на ftp сервере
func UploadFile(srcPath string, destPath string) {
	data, err := ioutil.ReadFile(srcPath)
	reader := bytes.NewBuffer(data)

	err = connection.Stor(ftpPath+destPath, reader)
	if err != nil {
		panic(err)
	}
}

// Запись фала по переданному пути из *[]byte
func WriteFile(path string, data *[]byte) {
	file, _ := os.Create(path)
	_, err := file.Write(*data)

	if err != nil {
		panic(err)
	}
}

// Загрузка файла с ftp сервера
func DownloadFile(srcPath, destPath string) {
	r, err := connection.Retr(srcPath)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	WriteFile(destPath, &buf)
}

// Удаление файла с удаленного сервера
func RemoveFile(path string) {
	connection.Delete(path)
}

// Инициализация
func Init() {
	connection.MakeDir(ftpPath)
	connection.ChangeDir(ftpPath)
}

// Вывод содержимого директории
func RetriveDir(path string) {
	walker := connection.Walk(path)
	for walker.Next() {
		fmt.Println(walker.Path())
	}
}

// Точка входа в программу
func main() {
	var err error
	ftpURLF := flag.String("url", "students.yss.su:21", "Server address")
	loginF := flag.String("user", "ftpiu8", "Login")
	passwordF := flag.String("pass", "3Ru7yOTA", "Password for login")

	flag.Parse()

	ftpURL = *ftpURLF
	login = *loginF
	password = *passwordF

	connection, err = ftp.Connect(ftpURL)
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Login(login, password)
	if err != nil {
		log.Fatal(err)
	}

	Init()
	UploadFile("README.txt", "/README.txt")
	DownloadFile("/tat/README.txt", "downloads/README.txt")
	MakeDir("AnotherOne")
	UploadFile("README.txt", "/AnotherOne/README.txt")
	RetriveDir(ftpPath)

	if err := connection.Quit(); err != nil {
		log.Fatal(err)
	}

}
