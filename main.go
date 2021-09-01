package main

import (
	"fmt"
	"inv/dig"
	"inv/up"
	"inv/whois"
	"os"

	"github.com/speedata/optionparser"
)

var (
	digVar   string
	upVar    string
	whoisVar string
	version  string
	banner   = "Usage: inv [OPTIONS] COMMAND\n"
)

func main() {
	op := optionparser.NewOptionParser()
	op.Banner = banner
	op.On("-u", "--up IP", "Check liveness stats", &upVar)
	op.On("-d", "--dig domain", "Dig DNS records", &digVar)
	op.On("-w", "--whois domain", "Whois domain information", &whoisVar)
	op.On("--version", "Print version")

	err := op.Parse()
	command := os.Args[1]
	if err != nil {
		fmt.Printf("Unknown option: %s\n\n", command)
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
	case "-u", "-up", "--up":
		up.Up(upVar)
	case "--version":
		fmt.Printf("Version: %s\n", version)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		op.Help()
		os.Exit(0)
	}
}
