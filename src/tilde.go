package main

import(
  "fmt"
  "os"
)

func main() {
  args := os.Args
  argv := args[1:]

  tilde = Tilde{false}

  if len(argv) > 1 {
    fmt.Println("Usage: tilde [script]")
    os.Exit(64)
  } else if len(argv) == 1 {
    tilde.RunFile(argv[0])
  } else {
    tilde.RunPrompt()
  }
}
