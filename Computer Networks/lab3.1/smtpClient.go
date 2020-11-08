package main

import (
	"bufio"
	"crypto/aes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/smtp"
	"os"
	"strconv"
)

type Mail struct {
	from, to, subject, body string
}

type Config struct {
	Username, Host string
	Port           int
	Pwd            []byte
}

type Intent struct {
	name         string
	inputReader  io.Reader
	outputWriter io.Writer
	handler      func(io.Reader, io.Writer)
}

func NewIntent(name string, input io.Reader, output io.Writer, handler func(io.Reader, io.Writer)) *Intent {
	return &Intent{name: name,
		inputReader:  input,
		outputWriter: output,
		handler:      handler,
	}
}

func (i Intent) Mathces(s string) bool {
	return i.name == s
}

var (
	cfg     *Config
	path    = flag.String("cfgPath", "./cfg.json", "Configuration Path")
	intents *[]Intent
	key     = []byte("passphrasewhichneedstobe32bytes!")
)

func parsecfg() {
	file, err := os.OpenFile(*path, os.O_RDONLY|os.O_CREATE, os.ModePerm.Perm())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg = &Config{}
	err = decoder.Decode(cfg)
	if err != nil {
		cfg = nil
	}
}

func stringToByte32(s string) []byte {
	out := make([]byte, 32)
	copy(out[0:], s)
	return out
}

func updatecfg(cfg *Config) {
	file, err := os.OpenFile(*path, os.O_WRONLY|os.O_CREATE, os.ModePerm.Perm())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmp := *cfg
	key := []byte("passphrasewhichneedstobe32bytes!")
	c, _ := aes.NewCipher(key)
	encryptedpwd := make([]byte, 32)

	c.Encrypt(encryptedpwd, cfg.Pwd)
	cfg = &Config{tmp.Username, tmp.Host, tmp.Port, encryptedpwd}

	data, err := json.Marshal(cfg)
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

func mainMenuIntent(input io.Reader, output io.Writer) {
	CreateIntent := func(name string, h func(io.Reader, io.Writer)) *Intent {
		return NewIntent(name,
			input,
			output,
			h)
	}
	intents := []*Intent{
		CreateIntent("send",
			func(input io.Reader, output io.Writer) {
				var to, subj, body string
				mail := Mail{}
				mail.from = cfg.Username
				fmt.Fprint(output, "Enter rec email addres: \n> ")
				fmt.Fscanf(input, "%s\n", &to)
				fmt.Fprint(output, "Enter mail subj: \n> ")
				fmt.Fscanf(input, "%s\n", &subj)
				fmt.Fprint(output, "Enter mail body: \n> ")
				fmt.Fscanf(input, "%s\n", &body)

				c, _ := aes.NewCipher(key)
				encryptedpwd := make([]byte, 32)
				c.Decrypt(encryptedpwd, cfg.Pwd)

				conn, err := tls.Dial("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port), tlsconfig)
				if err != nil {
					panic(err)
				}

				smtpClient, err = smtp.NewClient(conn, cfg.Host)
				if err != nil {
					panic(err)
				}

				auth := smtp.PlainAuth("", cfg.Username, string(encryptedpwd), cfg.Host)

				if err = smtpClient.Auth(auth); err != nil {
					panic(err)
				}

				if err = smtpClient.Rcpt(to); err != nil {
					panic(err)
				}

				w, _ := smtpClient.Data()
				fmt.Fprintf(w, "To:%s\r\nFrom:%s\r\nSubject:\r\n\r\ns%\r\n", to, cfg.Username, subj, body)
				defer smtpClient.Quit()
			}),
		CreateIntent("exit",
			func(input io.Reader, output io.Writer) {
				smtpClient.Close()
				os.Exit(0)
			}),
		CreateIntent("encryptpwd",
			func(input io.Reader, output io.Writer) {
				c, _ := aes.NewCipher(key)
				println(len(cfg.Pwd))
				res := make([]byte, 32)
				c.Decrypt(res, cfg.Pwd)
				fmt.Fprintf(output, "%s\n", string(res))
			}),
		CreateIntent("auth",
			func(input io.Reader, output io.Writer) {
				var host, username, pwd string
				var port int
				fmt.Fprint(output, "Enter smtp host\n> ")
				fmt.Fscanf(input, "%s\n", &host)
				fmt.Fprint(output, "Enter smtp port\n> ")
				fmt.Fscanf(input, "%d\n", &port)
				fmt.Fprint(output, "Enter smtp username\n> ")
				fmt.Fscanf(input, "%s\n", &username)
				fmt.Fprint(output, "Enter smtp password\n> ")
				fmt.Fscanf(input, "%s\n", &pwd)
				bpwd := make([]byte, 32)
				copy(bpwd, pwd)
				cfg = &Config{username, host, port, bpwd}
				updatecfg(cfg)
				//connect and auth
				tlsconfig = &tls.Config{
					InsecureSkipVerify: true,
					ServerName:         cfg.Host + ":" + strconv.Itoa(cfg.Port),
				}

				_, err := tls.Dial("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port), tlsconfig)
				if err != nil {
					panic(err)
				}

			})}
	fmt.Println("Welcome to awesome smtp client.")
	inputReader := bufio.NewReader(input)
	for {
	L:
		in, _, _ := inputReader.ReadLine()
		inputStr := string(in)

		for _, intent := range intents {
			if intent.Mathces(inputStr) {
				intent.handler(inputReader, output)
				goto L
			}
		}
		fmt.Fprintln(output, "No such command")
	}
}

var tlsconfig *tls.Config
var smtpClient *smtp.Client

func main() {
	flag.Parse()
	parsecfg()
	if cfg != nil {
		tlsconfig = &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         cfg.Host,
		}
	}
	mainMenuIntent(os.Stdin, os.Stdout)
}

//tarasovtest@ya.ru
//riznvqaairnjmntd
//
