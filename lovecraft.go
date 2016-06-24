package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/dborzov/lovecraft/world"
)

var w *world.World

func init() {
  w = world.NewWorld()
  err := w.Load("world")
  if err != nil {
    log.Fatalf("Failure loading universe: %s", err)
  }
}
func main() {
  reader := bufio.NewReader(os.Stdin)
  for true {
    fmt.Printf("\n>")
    text, _ := reader.ReadString('\n')
    command(text)
  }
}

func command(rawcmd string) {
  args := strings.Split(strings.TrimSuffix(rawcmd, "\n"), " ")
  if len(args) == 0 {
    return
  }
  cmd := args[0]
  switch cmd {
  case "help":
    fmt.Printf(`
         ls - to list elements in lands
         elements - to list available elements
         transforms - to list available transformations
         exit - to exit console
      `)
  case "ls":
    fmt.Printf("available lands\n")
    for i, val := range w.Lands {
      fmt.Printf("   %v --> %v\n", i, val.ElList)
    }
  case "elements":
    fmt.Printf("available elements\n")
    for key, val := range w.Elements {
      fmt.Printf("   %s: %v\n", key, val)
    }
  case "transforms":
    fmt.Printf("available transformation rules:\n")
    for key, val := range w.Transformations {
      fmt.Printf("   %s: %s\n", key, val)
    }
  case "exit":
    os.Exit(0)
  default:
    fmt.Printf("command \"%s\" is not recognized", cmd)
  }
}
