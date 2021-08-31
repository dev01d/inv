package up

import (
	"fmt"
	"net"
	"os"

	"github.com/go-ping/ping"
)

func pings(upVar string) {
	validateIP := net.ParseIP(upVar)
	if validateIP == nil {
		fmt.Printf("%s is not a valid IP Address\n", upVar)
		os.Exit(0)
	}
	ping, err := ping.NewPinger(upVar)
	if err != nil {
		fmt.Printf("Cannot resolve %s\n", upVar)
		os.Exit(0)
	}
	ping.Count = 4
	err = ping.Run()
	if err != nil {
		fmt.Printf("Cannot resolve %s\n", upVar)
		os.Exit(0)
	}
	results := ping.Statistics()
	fmt.Printf("\nResults of ping -c 4")
	fmt.Printf("\nHost up on IP:\t\t%s\nAverage latency:\t%s\nAverage packet loss:\t%.2f\n\n",
		results.IPAddr, results.AvgRtt, results.PacketLoss)
}

func Up(upVar string) {
	pings(upVar)
}
