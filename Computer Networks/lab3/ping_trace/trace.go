package main

import (
	"fmt"
	"flag"
	"github.com/lixiangzhong/traceroute"
)

var (
	url     = flag.String("url", "ya.ru", "Url")
)

func main() {
	flag.Parse()
	t := traceroute.New(*url)

	result, err := t.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}
}