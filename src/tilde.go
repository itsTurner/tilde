package main

import(
  "fmt"
  "os"
  "io/ioutil"
  "bufio"
)

type Tilde struct {
  HadError bool
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var tilde Tilde

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


func (t Tilde) Error(line int, message string) {
  t.Report(line, "", message)
}

func (t Tilde) Report(line int, where string, message string) {
  fmt.Println("[Line %i] Error%s: %s", line, where, message)
  t.HadError = true
}

func (t Tilde) RunFile(path string) {
  dat, err := ioutil.ReadFile(path)
  check(err)
  data := string(dat[:])
  t.Run(data)

  if t.HadError { os.Exit(65) }
}

func (t Tilde) RunPrompt() {
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    t.Run(text)
    t.HadError = false
  }
}

func (t Tilde) Run(source string) {
  scanner := Scanner{source}
  tokens := scanner.ScanTokens()

  for _, token := range(tokens) {
    fmt.Println(token)
  }
}
