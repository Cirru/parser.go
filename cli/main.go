
package main

import (
  "io/ioutil"
  "github.com/Cirru/cirru-parser.go"
)

func main() {
  b, err := ioutil.ReadFile("../cirru/spaces.cirru")
  if err != nil {
    panic(err)
  }
  parser := cirru.NewParser()
  for _, c := range b {
    parser.Read(rune(c))
  }
  parser.GetAst()
}