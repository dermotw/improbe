package Ping

import (
	"log"
	"os/exec"
	"regexp"
	"encoding/json"
	"time"
	"strconv"
)

// Ping function
//
func Ping( host string ) (b string) {
	// Cheat and call fping, the command looks like this:
	//
	// fping -p 20 -b 1400 -a -C 20 <host>
	//
	// -p 20 					20 milliseconds between each Ping
	// -b 1400				1400 byte payload
	// -a							only report hosts that are alive
	// -C 20					20 echo requests, formatted in a way that's easy to work
	//								with programmatically
	//
	out, err := exec.Command("fping", "-p 20", "-b 1400", "-a", "-C 20", host).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)

	// Regex for extracting the results - we only want the round-trip times, so
	// match anything that looks like x.xx *or* the packets that timeout (-)
	//
	// Note that IP addresses will match, which is why the entries in host.list
	// **must** be hostnames!
	//
	re := regexp.MustCompile(`\d+\.\d{1,2}|-`)

	// theMatches will store the matches!
	//
	theMatches := re.FindAllStringSubmatch( s, -1 )

	// Add the hostname to the end of the array so that we know what it is when
	// we pass the results to the remote server
	//
	results := append( theMatches, []string{host} )
	results = append( results, []string{strconv.FormatInt( time.Now().Unix(), 10 )} )

	// Convert the results array to a JSON object to make it easier to work with
	//
	j, err := json.Marshal( results )
	b = string(j)
	if err != nil {
		log.Fatal( err )
	}
	return

}
