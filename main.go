package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, _ := net.LookupMX(domain)
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, _ := net.LookupTXT(domain)
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("domain--%v,hasMX-- %v, hasSPF--%v,sprRecord-- %v, hasDMARC--%v, dmarcRecord--%v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

}
