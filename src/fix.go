package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
  "bytes"
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
    var buffer bytes.Buffer
    buffer.WriteString(scanner.Text())
    buffer.WriteString("\n")
    fmt.Fprintf(outFile, buffer.String())
  }

  // Close files
  inFile.Close()
  outFile.Close()
}
