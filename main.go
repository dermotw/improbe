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

func main() {
  file, err := os.Open( "./hosts.list" )
  if err != nil {
    log.Fatal( err )
  }
  defer file.Close()

  scanner := bufio.NewScanner( file )
  for scanner.Scan() {
    host := scanner.Text()
    r := Ping.Ping( host )
    req, err := http.NewRequest( "POST", "http://10.10.1.106/improbe/results", bytes.NewBuffer( []byte(r) ) )
    req.Header.Set( "X-Custom-Header", "blah" )
    req.Header.Set( "Content-Type", "application/json" )

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
      log.Fatal( err )
    }

    defer res.Body.Close()

    fmt.Printf( "%s\n", res.Status )
  }
}
