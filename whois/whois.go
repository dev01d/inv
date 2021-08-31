package whois

import (
	"fmt"
	"net"

	whois "github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func Whois(whoisVar string) {
	result, err := whois.Whois(whoisVar)
	if err != nil {
		fmt.Println(err)
	} else {
		data, err := whoisparser.Parse(result)
		isAddr := net.ParseIP(whoisVar)
		if isAddr != nil {
			fmt.Println("Only domains please\n \bThat query returns:", data)
		} else {
			if err == nil {
				fmt.Println("\n############ [WHOIS QUERY] #############")
				fmt.Println("Domain Name:", data.Domain.Domain)
				fmt.Println("Exp Date:", data.Domain.ExpirationDate)
				fmt.Println("Updated:", data.Domain.UpdatedDate)
				fmt.Println("Created:", data.Domain.CreatedDate)
				fmt.Println("Status:", data.Domain.Status[0])
				fmt.Println("DNSSec:", data.Domain.DNSSec)
				fmt.Println("Reg Name:", data.Registrar.Name)
				nsArr := data.Domain.NameServers
				for i, ns := range nsArr {
					fmt.Println("Name Server", i+1, "\b:", ns)
				}
				fmt.Printf("########################################\n\n")
			} else {
				fmt.Println("\nDoamin not found or not registered")
			}
		}
	}
}
