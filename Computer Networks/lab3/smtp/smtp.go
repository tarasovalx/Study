package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/smtp"
)

var addr, subject, message string

var
(
	host = flag.String("host", "smtp.yandex.ru" , "host")
	sender = flag.String("user", "****@yandex.com" , "Url")
)

func main() {
	flag.Parse()
	auth := smtp.PlainAuth("", *sender, "****", *host)

	fmt.Print("Enter reciever adress: ")
	if _, err := fmt.Scan(&addr); err != nil {
		panic(err)
	}

	fmt.Print("Enter Subject: ")
	if _, err := fmt.Scan(&subject); err != nil {
		panic(err)
	}

	fmt.Print("Enter message: ")
	if _, err := fmt.Scan(&message); err != nil {
		panic(err)
	}

	msg := []byte("To: " + addr + "\r\n" +
		"From: " + *sender + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: *host,
	}

	conn, err := tls.Dial("tcp", *host+":465", tlsConfig)

	client, err := smtp.NewClient(conn, *host)
	if err != nil{
		log.Panic(err)
	}

	if err = client.Auth(auth); err != nil{
		log.Panic(err)
	}

	if err = client.Mail(*sender); err != nil{
		log.Panic(err)
	}

	if err = client.Rcpt(addr); err != nil{
		log.Panic(err)
	}

	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = wc.Write(msg)
	if err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
