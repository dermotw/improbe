package main

import (
  "improbe/Ping"
  "bufio"
  "log"
  "os"
  "fmt"
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
    fmt.Printf( "%s\n", r )
  }
}
