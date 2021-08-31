package dig

import (
	"fmt"
	"net"
)

var handleErr = "no record found\n"

func aRecord(digVar string) {
	names, err := net.LookupIP(digVar)
	if err != nil {
		fmt.Print(handleErr)
	}
	if len(names) == 0 {
		fmt.Print(handleErr)
	}
	fmt.Printf("\n----  A RECORD(S)  ----\n")
	for _, name := range names {
		fmt.Printf(" %s\n", name)
	}
}

func mxRecord(digVar string) {
	names, err := net.LookupMX(digVar)
	if err != nil {
		fmt.Print(handleErr)
	}
	if len(names) == 0 {
		fmt.Print(handleErr)
	}
	fmt.Printf("\n----  MX RECORD(S)  ----\n")
	for _, name := range names {
		fmt.Printf(" %s\n", name.Host)
	}
}

func nsRecord(digVar string) {
	names, err := net.LookupNS(digVar)
	if err != nil {
		fmt.Print(handleErr)
	}
	if len(names) == 0 {
		fmt.Print(handleErr)
	}
	fmt.Printf("\n----  NS RECORD(S)  ----\n")
	for _, name := range names {
		fmt.Printf(" %s\n", name.Host)
	}
}

func txtRecord(digVar string) {
	names, err := net.LookupTXT(digVar)
	if err != nil {
		fmt.Print(handleErr)
	}
	if len(names) == 0 {
		fmt.Print(handleErr)
	}
	fmt.Printf("\n---- TXT RECORD(S) ----\n")
	for _, name := range names {
		fmt.Printf(" %s\n", name)
	}
	fmt.Printf("\n")
}

func Dig(digVar string) {
	aRecord(digVar)
	mxRecord(digVar)
	nsRecord(digVar)
	txtRecord(digVar)
}
