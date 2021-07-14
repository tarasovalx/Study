package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type sshConfig struct {
	User, Password, Host string
}

type configuration struct {
	Hosts []sshConfig
}

type cmdTask struct {
	cmd  string
	host net.Addr
}

var (
	cfg     configuration
	clients []*ssh.Client
	channel chan cmdTask
)

func parseConfig() {
	file, err := os.Open("./cfg.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg = configuration{}
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Welcome to multi ssh client")
	flag.Parse()

	channel = make(chan cmdTask)
	timeout := time.After(25 * time.Second)
	parseConfig()
	fmt.Println(cfg)

	clients = make([]*ssh.Client, len(cfg.Hosts))
	for i, h := range cfg.Hosts {
		client, e := ssh.Dial("tcp", h.Host, &ssh.ClientConfig{
			User: h.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(h.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
		if e != nil {
			log.Println(e.Error())
		}
		clients[i] = client
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := reader.ReadString('\n')
		if strings.TrimSpace(cmd) == "" {
			continue
		}
		for _, c := range clients {
			go executeCommand(strings.TrimSpace(cmd), c)
		}
		for i := len(clients); i != 0; {
			select {
			case res := <-channel:
				fmt.Printf("hostname: %s\n", res.host)
				fmt.Println(res.cmd)
				i--

			case <-timeout:
				fmt.Println("Execution timeout...")
				break
			}
		}
		if cmd == "exit" {
			return
		}
	}
}

func executeCommand(cmd string, c *ssh.Client) {
	start := time.Now()
	session, e := c.NewSession()
	if e != nil {
		log.Fatal(e)
	}

	o, e := session.Output(cmd)
	channel <- cmdTask{host: c.Conn.RemoteAddr(), cmd: string(o)}
	fmt.Printf("Estimated time %f\n", time.Now().Sub(start).Seconds())
	defer session.Close()
}
