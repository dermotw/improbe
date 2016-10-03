package Ping

import (
//	"fmt"
	"log"
	"os/exec"
	"os"
	"regexp"
	"encoding/json"
)

func Ping( host string ) {
	out, err := exec.Command("fping", "-p 20", "-b 1400", "-a", "-C 20", host).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	re := regexp.MustCompile(`\d+\.\d{1,2}|-`)

	theMatches := re.FindAllStringSubmatch( s, -1 )

	b, err := json.Marshal( theMatches )
	if err != nil {
		log.Fatal( err )
	}
	os.Stdout.Write( b )

}
