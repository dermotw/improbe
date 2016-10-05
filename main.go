package main

import (
  "improbe/Ping"
  "bufio"
  "log"
  "os"
  "fmt"
  "net/http"
  "net/url"
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
    theData := url.Values{}
    theData.Add( "result", r )
    res, err := http.PostForm( "http://10.10.1.106/improbe/results.php", theData )

    if err != nil {
      log.Fatal( err )
    }
    defer res.Body.Close()

    fmt.Printf( "%s\n", res.Status )
  }
}
