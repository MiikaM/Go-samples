package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	maxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(maxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v:=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v:=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("\nDomain: %v\nHas MX: %v\nHas SPF: %v\nSPF Record: %v\nHas DMARC: %v\nDMARC Record: %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
