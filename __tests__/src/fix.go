package src

import (
  "os"
  "fmt"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Argument path missing.")
  }

}
