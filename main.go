package main

import (
  "improbe/Ping"
  "bufio"
  "log"
  "os"
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
    Ping.Ping( host )
  }
}
