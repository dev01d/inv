package up

import (
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/go-ping/ping"
)

func pings(upVar string) {
	validateIP := net.ParseIP(upVar)
	if validateIP == nil {
		color.Red("%s is not a valid IP Address\n", upVar)
		os.Exit(0)
	}
	ping, err := ping.NewPinger(upVar)
	if err != nil {
		color.Red("Cannot resolve %s\n", upVar)
		os.Exit(0)
	}
	ping.Count = 4
	ping.Timeout = 3 * time.Second
	err = ping.Run()
	if err != nil {
		color.Red("Cannot resolve %s\n", upVar)
		os.Exit(0)
	}
	results := ping.Statistics()
	color.Blue("\nResults of ping -c 4")
	fmt.Printf("Host up on IP:\t\t%s\nAverage latency:\t%s\nAverage packet loss:\t%.2f\n\n",
		results.IPAddr, results.AvgRtt, results.PacketLoss)
}

func Up(upVar string) {
	pings(upVar)
}
