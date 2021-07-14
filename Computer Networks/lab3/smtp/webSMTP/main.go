package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "auth.html")
	})

	http.HandleFunc("/auth", AuthHandler)
	http.HandleFunc("/client/", ClientHandler)
	http.HandleFunc("/send", SendHandler)

	if err := http.ListenAndServe(":8005", nil); err != nil {
		log.Panic(err)
	}
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println("AUTH HANDLER\n", r.Form)

	jsonLogin, _ := json.Marshal(struct {
		Login string
		Pwd   string
	}{Login: r.Form["login"][0], Pwd: r.Form["password"][0]})

	http.SetCookie(w, &http.Cookie{
		Name:    "loginfo",
		Value:   base64.StdEncoding.EncodeToString(jsonLogin),
		Expires: time.Now().Add(30 * time.Minute),
	})

	http.Redirect(w, r, "/client/", 301)
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println("CLIENT HANDLER")

	path := strings.Replace(r.URL.Path, "/client", "", 1)
	if len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	//c := MailLogin(w, r)
	//defer c.Quit()

	http.ServeFile(w, r, "index.html")
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("SEND HANDLER")

	//c := MailLogin(w, r)

	host, port, _ := net.SplitHostPort("smtp.mail.ru:465")
	cookie, err := r.Cookie("loginfo")
	if err != nil {
		log.Panic(err)
	}
	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Panic(err)
	}
	info := struct {
		Login string
		Pwd   string
	}{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		log.Panic(err)
	}
	auth := smtp.PlainAuth("", fmt.Sprintf("%s@mail.ru", info.Login), info.Pwd, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", host, port), tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		http.Redirect(w, r, "/auth", 401)
	}
	defer c.Quit()
	if err = c.Auth(auth); err != nil {
		http.Redirect(w, r, "/auth", http.StatusUnauthorized)
	}
	log.Println("--LOGIN OK--")

	if err != nil {
		log.Panic(err)
	}

	from := mail.Address{
		Name:    r.Form["yourName"][0],
		Address: fmt.Sprintf("%s@mail.ru", info.Login),
	}

	to := mail.Address{
		Name:    r.Form["recvName"][0],
		Address: r.Form["to-inp"][0],
	}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = r.Form["subject"][0]
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + r.Form["text-inp"][0]

	if err = c.Mail(from.Address); err != nil {
		log.Println("---------------", from.Address)
		log.Panic(err)
	}
	if err = c.Rcpt(r.Form["to-inp"][0]); err != nil {
		log.Panic(err)
	}
	wr, err := c.Data()
	defer wr.Close()
	if err != nil {
		log.Panic(err)
	}
	_, err = wr.Write([]byte(message))
	http.ServeFile(w, r, "success.html")
}

func MailLogin(w http.ResponseWriter, r *http.Request) *smtp.Client {
	host, port, _ := net.SplitHostPort("smtp.mail.ru:465")
	cookie, err := r.Cookie("loginfo")
	if err != nil {
		log.Panic(err)
	}
	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Panic(err)
	}
	info := struct {
		login string
		pwd   string
	}{}
	err = json.Unmarshal(data, &info)
	log.Println(info)
	if err != nil {
		log.Panic(err)
	}
	auth := smtp.PlainAuth("", fmt.Sprintf("%s@mail.ru", info.login), info.pwd, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", host, port), tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		http.Redirect(w, r, "/auth", 401)
		return nil
	}
	if err = c.Auth(auth); err != nil {
		http.Redirect(w, r, "/auth", http.StatusUnauthorized)
	}
	log.Println("--LOGIN OK--")
	return c
}
