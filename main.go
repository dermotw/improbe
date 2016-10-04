package main

import (
  "bytes"
  "improbe/Ping"
  "bufio"
  "log"
  "os"
  "fmt"
  "net/http"
)

// main function
//
func main() {
  // Open hosts.list, a list of hosts to ping - one per line
  //
  file, err := os.Open( "./hosts.list" )
  if err != nil {
    log.Fatal( err )
  }
  defer file.Close()

  // Create a bufio NewScanner object to read in the content of
  // hosts.list, one line at a time
  //
  scanner := bufio.NewScanner( file )
  for scanner.Scan() {
    host := scanner.Text()

    // Send the IP address to our Ping function
    //
    r := Ping.Ping( host )

    // Set up an HTTP Request object and stuff
    //
    req, err := http.NewRequest( "POST", "http://10.10.1.106/improbe/results", bytes.NewBuffer( []byte(r) ) )
    req.Header.Set( "X-Custom-Header", "blah" )
    req.Header.Set( "Content-Type", "application/json" )

    // HTTP client
    //
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
      log.Fatal( err )
    }
    defer res.Body.Close()

    fmt.Printf( "%s\n", res.Status )
  }
}
