package main

import (
	"fmt"
	"inv/cmd/dig"
	"inv/cmd/ping"
	"inv/cmd/whois"
	"os"

	"github.com/fatih/color"
	"github.com/speedata/optionparser"
)

var (
	digVar   string
	pingVar  string
	whoisVar string
	version  string
	banner   = "Usage: inv [OPTIONS] COMMAND\n"
)

func main() {
	op := optionparser.NewOptionParser()
	op.Banner = banner
	op.On("-d", "--dig domain", "Dig DNS records", &digVar)
	op.On("-p", "--ping IP", "Check liveness stats", &pingVar)
	op.On("-w", "--whois domain", "Whois domain information", &whoisVar)
	op.On("--version", "Print version")

	err := op.Parse()
	if err != nil {
		color.Red("Unknown option: %s\n", os.Args[1])
		op.Help()
		os.Exit(0)
	} else if len(os.Args) < 2 {
		op.Help()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "-d", "-dig", "--dig":
		dig.Dig(digVar)
	case "-w", "-whois", "--whois":
		whois.Whois(whoisVar)
	case "-p", "-ping", "--ping":
		ping.Ping(pingVar)
	case "--version":
		fmt.Printf("Version: %s\n", version)
	default:
		color.Red("Unknown command: %s\n", os.Args[1])
		op.Help()
		os.Exit(0)
	}
}
