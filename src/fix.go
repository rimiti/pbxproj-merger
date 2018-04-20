package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Argument path missing.")
  }

  // Input file
  inFile, _ := os.Open(os.Args[1])
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  // Output file
  outFile, err := os.Create("output.txt")
  if err != nil {
    log.Fatal("Cannot create file", err)
  }
  defer outFile.Close()

  for scanner.Scan() {
    fmt.Fprintf(outFile, scanner.Text(), "\n")
    // fmt.Println("=====>", scanner.Text())
  }
}
