package ping

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-ping/ping"
)

func pings(pingVar string) {
	validateIP := net.ParseIP(pingVar)
	if validateIP == nil {
		color.Red("%s is not a valid IP Address\n", pingVar)
		os.Exit(0)
	}
	ping, err := ping.NewPinger(pingVar)
	if err != nil {
		color.Red("Cannot ping host %s\n", pingVar)
		os.Exit(0)
	}
	ping.Count = 4
	ping.Timeout = 3 * time.Second
	err = ping.Run()
	if err != nil {
		color.Red("Cannot ping host %s\n", pingVar)
		os.Exit(0)
	}
	results := ping.Statistics()
	if results.PacketLoss > 99 {
		color.Red("Cannot ping host %s\n", pingVar)
		os.Exit(0)
	}
	color.Cyan("\nResults of ping -c 4")
	fmt.Printf("Host up on IP:\t\t%s\nAverage latency:\t%s\nAverage packet loss:\t%.2f\n\n",
		results.IPAddr, results.AvgRtt, results.PacketLoss)
}

func Ping(pingVar string) {
	pings(pingVar)
}
