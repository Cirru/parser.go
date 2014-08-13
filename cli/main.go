
package main

import (
  "io/ioutil"
  "github.com/Cirru/cirru-parser.go"
  "fmt"
  "encoding/json"
)

func main() {
  samples := []string{"quote",
    "comma",
    "folding",
    "indent",
    "line",
    "parentheses",
    "spaces",
    "unfolding",
    "demo",
    "html",
  }

  for _, sample := range(samples) {
    cirruFile := fmt.Sprintf("../cirru/%s.cirru", sample)
    jsonFile := fmt.Sprintf("../ast/%s.json", sample)
    b, _ := ioutil.ReadFile(cirruFile)
    b2, _ := ioutil.ReadFile(jsonFile)
    parser := cirru.NewParser()
    for _, c := range b {
      parser.Read(rune(c))
    }
    parser.Complete()

    content, _ := json.MarshalIndent(parser.ToJSON(), "", "  ")
    gotAst := string(content)
    wantedAst := string(b2)

    if gotAst != wantedAst {
      println(sample, "-- not matching, break:")
      println(gotAst)
      break
    } else {
      println(sample, "-- matches")
    }

  }
}