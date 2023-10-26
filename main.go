// package dnslookup provides tools to obtain
// A, CNAME, MS, NS, PTR, TXT records of a given domain
package dnslookup

import (
	"net"
	"sync"
)

// mutex is a sync.Mutex object created to be used to prevent
// problems that may arise due to concurrency during DNS lookup.
var mutex sync.Mutex

// DnsRecord represents a DNS record for a specific domain
type DnsRecord struct {
	// domain , the domain for which this DNS record is stored
	domain string

	// aRecords , a slice of A (IPv4) records.
	aRecords []net.IP

	// cnameRecords, a string of CNAME (Canonical Name) records.
	cnameRecords string

	// mxRecords, a slice of MX (Mail Exchange) records.
	mxRecords []net.MX

	// nsRecords, a slice of NS (Name Server) records.
	nsRecords []net.NS

	// ptrRecords, a slice of PTR (Pointer) records.
	ptrRecords []string

	// txtRecords, a slice of TXT (Text) records.
	txtRecords []string
}

// NewDnsRecord is a constructor function simplifies the process of
// creating a new 'DnsRecord' instance by encapsulating the initialization logic
//
// domainName: a string representing of the domain name associated with
// the DNS record.
//
// NewDnsRecord returns a *DnsRecord object
func NewDnsRecord(domainName string) *DnsRecord {
	return &DnsRecord{
		domain: domainName,
	}
}

// GetARecords method is designed to retrieve the A
// records for the given domain. It caches the results in
// the 'DnsRecord' instance to avoid unnecessary DNS lookups
// and uses a mutex to ensure safe concurrent access to the
// cached data.
//
// If the A records are already cached, it returns the cached
// records; otherwise, it performs a DNS lookup to fetch the
// records and caches them for future use.
//
// It returns a slice of that domain's IP addresses in the DnsRecords instance
func (d *DnsRecord) GetARecords() []net.IP {
	if d.aRecords == nil {
		aRecords, err := net.LookupIP(d.domain)
		if err == nil {
			mutex.Lock()
			d.aRecords = append(d.aRecords, aRecords...)
			mutex.Unlock()
		}
		return d.aRecords
	}
	return d.aRecords
}

// GetCnameRecords method is designed to retrieve the CNAME
// record for the given domain. It caches the result in the
// DnsRecord instance to avoid unnecesarry DNS lookups and uses
// a mutex to ensure safe concurrent access to the cached data.
//
// If the CNAME record is already cached, it returns the cached
// record; otherwise, it performs a DNS lookup to fetch the records
// and caches them for future use.
//
// It returns a string of Canonical Name Record belongs to that domain
// in the DnsRecord instance
func (d *DnsRecord) GetCnameRecords() string {
	if d.cnameRecords == "" {
		cnameRecords, err := net.LookupCNAME(d.domain)
		if err == nil {
			mutex.Lock()
			d.cnameRecords = cnameRecords
		}
		mutex.Unlock()
		return d.cnameRecords
	}
	return d.cnameRecords
}

// GetMxRecords method is designed to retrieve MX records for the
// given domain. It caches the result in the DnsRecords instance to
// avoid unnecessary DNS lookups and uses a mutex to ensure safe concurrent
// access to the cached data.
//
// If the MX record is already cached, it returns the cached
// record; otherwise, it performs a DNS lookup to fetch the records
// and caches them for future use.
//
// It returns a slice of mail server names belongs to that domain in the
// DnsRecord instance
func (d *DnsRecord) GetMxRecords() []net.MX {
	if d.mxRecords == nil {
		mxRecords, err := net.LookupMX(d.domain)
		if err == nil {
			mutex.Lock()
			d.mxRecords = make([]net.MX, len(mxRecords))
			for i, record := range mxRecords {
				d.mxRecords[i] = *record
			}
			mutex.Unlock()
		}
		return d.mxRecords
	}
	return d.mxRecords
}

// GetNsRecords method is designed to retrieve NS records for the
// given domain. It caches the result in the DnsRecord instance to avoid
// unnecessary DNS lookups and uses a mutex to ensure safe concurrent access
// to the cached data.
//
// If NS records are already cached, it returns the cached record;
// otherwise, it performs a DNS lookup to fetch the records and caches
// them for future use
//
// It returns a slice of Name Server addresses belong to that domain
// in the DnsRecord instance
func (d *DnsRecord) GetNsRecords() []net.NS {
	if d.nsRecords == nil {
		nsRecords, err := net.LookupNS(d.domain)
		if err == nil {
			mutex.Lock()
			d.nsRecords = make([]net.NS, len(nsRecords))
			for i, record := range nsRecords {
				d.nsRecords[i] = *record
			}
			mutex.Unlock()
		}
	}
	return d.nsRecords
}

// GetPtrRecords method is designed to retrieve PTR records for the
// given domain. It caches the result in the DnsRecord instance to avoid
// unnecessary DNS lookups and uses a mutex to ensure safe concurrent access
// to the cached data.
//
// If PTR records are already cached , it returns the cached records;
// otherwise, it performs a DNS lookup to fetch the records and caches them
// for future use.
//
// It returns a slice of strings of Pointer records belong to that domain
// in the DnsRecord instance
func (d *DnsRecord) GetPtrRecords() []string {
	if d.ptrRecords == nil {
		if d.aRecords == nil {
			d.GetARecords()
			for _, v := range d.aRecords {
				ptr, _ := net.LookupAddr(v.String())
				d.ptrRecords = append(d.ptrRecords, ptr...)
			}
		}
		for _, v := range d.aRecords {
			ptr, _ := net.LookupAddr(v.String())
			d.ptrRecords = append(d.ptrRecords, ptr...)
		}
	}
	return d.ptrRecords
}

// GetTxtRecords method is designed to retrieve TXT records for the given
// domain. It caches the result in the DnsRecords instance to avoid unnecessary
// DNS lookups and uses a mutex to ensure safe concurrent access to the cached data
//
// If TXT records are already cached, it returns the cached records; otherwise;
// it performs a DNS lookup to fetch the records and caches them for future use
//
// It returns a slice of strings of TXT records belong to that domain
// in DnsRecord instance
func (d *DnsRecord) GetTxtRecords() []string {
	if d.txtRecords == nil {
		txtRecords, err := net.LookupTXT(d.domain)
		if err == nil {
			mutex.Lock()
			d.txtRecords = txtRecords
		}
		mutex.Unlock()
	}
	return d.txtRecords
}

// GetAllRecords method retrieves and collect various DNS records
// for the given domain and returns them in a structured map. Each
// type of DNS record (A,CNAME,MX,NS,PTR,TXT) is obtained by calling
// separate methods, and the results are stored in the map under
// corresponding keys.
//
// It returns a map where the keys are strings and the values associated
// with these keys can be of any data type. This allows the function to store
// various types of DNS records in the same map.
func (d *DnsRecord) GetAllRecords() map[string]interface{} {

	var records = make(map[string]interface{})

	records["A records"] = d.GetARecords()
	records["CNAME records"] = d.GetCnameRecords()
	records["MX records"] = d.GetMxRecords()
	records["NS records"] = d.GetNsRecords()
	records["PTR records"] = d.GetPtrRecords()
	records["TXT records"] = d.GetTxtRecords()

	return records
}
