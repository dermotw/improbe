package Ping

import (
//	"fmt"
	"log"
	"os/exec"
	"regexp"
	"encoding/json"
)

func Ping( host string ) (b string) {
	out, err := exec.Command("fping", "-p 20", "-b 1400", "-a", "-C 20", host).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	re := regexp.MustCompile(`\d+\.\d{1,2}|-`)

	theMatches := re.FindAllStringSubmatch( s, -1 )

	results := append( theMatches, []string{host} )

	j, err := json.Marshal( results )
	b = string(j)
	if err != nil {
		log.Fatal( err )
	}
	return

}
