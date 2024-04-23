package nmap

import (
	"fmt"
	"log"

	"github.com/Ullaakut/nmap/v2"
	"github.com/fatih/color"
)

func Nmap(nmapVar string) {
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(nmapVar),
		nmap.WithFastMode(),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()

	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host")
		color.Cyan(" %q", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}
