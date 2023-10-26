# dnslookup

The dnslookup package provides tools to retrieve DNS records of a given domain.

It has mutex feature in each method so that it prevents problems that may occur related to concurrency.

## Installation
Install using the “go get” command:
```bash
go get -u github.com/yagizhanbilaldurak/dnslookup
```
## Usage
Create a new instance with NewDnsRecord :
```go
domain := "example.com"
newrecord := NewDnsRecord(domain)
```
Depending on need, for each DNS record, separate methods can be used.
```go
//for A (IPv4) records
newrecord.GetARecords()

//for CNAME (Canonical Name) records
newrecord.GetCnameRecords()

//for MX (Mail Exchange) records
newrecord.GetMxRecords()

//for NS (Name Server) records
newrecord.GetNsRecords()

//for PTR (Pointer) records
newrecord.GetPtrRecords()

//for TXT (Text) records
newrecords.GetTxtRecords()
If all records (A, CNAME, MX, NS, PTR, TXT) are needed, use simply :

newrecord.GetAllRecords()
//it will return a structured map with string keys and interface{} values.
//The returned map, contains DNS records organized by their record types as keys.
```
## Example Program
With this program, dnslookup package can be used with parameters.
```go
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
```
## Usage
Firstly, program must be built:
```bash
go build main.go
```
Now, run the program with command:

Note: if desired, search type can be change into “a, cname, mx, ns, ptr, txt”.

for a,cname,mx,ns,ptr,txt records; "all" parameter can be used
```bash
go run main.go -domain example.net -s all

```
Output:
```bash
map[A records:[93.184.216.34] CNAME records:example.net. MX records:[{. 0}] NS records:[{a.iana-servers.net.} {b.iana-servers.net.}] PTR records:[] TXT records:[v=spf1 -all 4wgz0ccyj83cx2y6xfpmmrp6w2d8gv2v]]
```
