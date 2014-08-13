
package main

import (
  "io/ioutil"
  "github.com/Cirru/cirru-parser.go"
  "fmt"
  "encoding/json"
)

func main() {
  b, err := ioutil.ReadFile("../cirru/comma.cirru")
  if err != nil {
    panic(err)
  }
  parser := cirru.NewParser()
  for _, c := range b {
    parser.Read(rune(c))
  }
  parser.Complete()

  content, err := json.MarshalIndent(parser.ToJSON(), "", "  ")
  if err != nil {
    fmt.Printf("error:", err)
  }
  println(string(content))
}