package main

import (
	"flag"
	"fmt"

	"github.com/yagizhanbilaldurak/dnslookup"
)

func main() {
	// Define and parse command-line flags
	domainPtr := flag.String("domain", "", "domain")
	searchTypePtr := flag.String("s", "", "searchType")
	flag.Parse()

	// Retrieve values from command-line flags
	domain := *domainPtr
	searchType := *searchTypePtr

	// Create a DNS record instance
	recorder := dnslookup.NewDnsRecord(domain)

	// Check if the searchType is valid
	if !isValidSearchType(searchType) {
		fmt.Println("error: -domain and -s parameters required. usage: dnssearch -domain example.net -s all")
		return
	} else {
		// Based on the searchType, perform DNS record lookup and print the results
		switch searchType {
		case "a":
			fmt.Println(recorder.GetARecords())
		case "all":
			fmt.Println(recorder.GetAllRecords())
		case "cname":
			fmt.Println(recorder.GetCnameRecords())
		case "mx":
			fmt.Println(recorder.GetMxRecords())
		case "ns":
			fmt.Println(recorder.GetNsRecords())
		case "ptr":
			fmt.Println(recorder.GetPtrRecords())
		case "txt":
			fmt.Println(recorder.GetTxtRecords())
		}
	}

}

func isValidSearchType(searchType string) bool {

	searchOptions := []string{"a", "all", "cname", "mx", "ns", "ptr", "txt"}

	found := true

	for _, v := range searchOptions {
		if searchType != v {
			found = false
		} else {
			return true
		}
	}

	return found

}

