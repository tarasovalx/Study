package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"log"
	"runtime"
	"flag"
	"time"
)

type chs chan string

var (
	url     = flag.String("url", "ya.ru", "Url")
	iters = flag.Int("iters", 5, "iterations")
)

func main() {
	flag.Parse()
	starttime := time.Now()
	runtime.GOMAXPROCS(2)
	result := make(chs)
	for i := 0; i < *iters; i++ {
		// go func(i int) {
			fmt.Println(i)
			result.StartPinging(i, *url)
		// }(i)
	}
	for i := 0; i < *iters; i++ {
		select {
		case res := <-result:
			fmt.Println(res)
		case <-time.After(10 * time.Second):
			fmt.Println("Timed out")
			return
		}
	}
	t := time.Now()
	elapsed := t.Sub(starttime)
	fmt.Println(elapsed.String())
}

func (result *chs) StartPinging(i int, addr string) {
	fmt.Println("pinging...")
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		log.Println(err)
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% 				packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = 1
	pinger.SetPrivileged(true)
	pinger.Run()
	*result <- fmt.Sprint(i, ": ", pinger.Statistics())
}
