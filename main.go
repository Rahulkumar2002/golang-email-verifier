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
	fmt.Println("domain , hasMX , hasSPF , spfRecord , hasDMARC , dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while in main: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasDMARC, hasMX, hasSPF bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error while in mx: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	textRecord, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error while in spf: %v\n", err)
	}

	for _, record := range textRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error while in dmarc: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v ,%v , %v ,%v ,%v ,%v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
