package main

import (
  "fmt"
  "os"
  "io/ioutil"
  tokens "format/tokenize"
  color "format/colorize"
)

func main(){
  args := os.Args
  if len(args) != 2 {
    fmt.Println("Usage: go run format.go <json file>")
    os.Exit(1)
  }

  filename := args[1]

  b, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Error reading file:", err)
    os.Exit(1)
  }

  res := tokens.Tokenize(string(b))
  // fmt.Println(res)
  color.ColorizeTokens(&res)
}
