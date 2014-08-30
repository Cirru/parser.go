
package parser

import (
  "testing"
  "io/ioutil"
  "encoding/json"
  "fmt"
)

func TestParser(t *testing.T) {
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
    cirruFile := fmt.Sprintf("./cirru/%s.cirru", sample)
    jsonFile := fmt.Sprintf("./json/%s.json", sample)
    b, _ := ioutil.ReadFile(cirruFile)
    b2, _ := ioutil.ReadFile(jsonFile)

    gotAst := ExampleNewParser(b)
    wantedAst := string(b2)

    if gotAst != wantedAst {
      println(sample, "\t-- not matching, break:")
      println(gotAst)
      return
    } else {
      println(sample, "\t-- matches")
    }

  }
}

func ExampleNewParser(b []byte) string {
  p := NewParser()
  for _, c := range b {
    p.Read(rune(c))
  }
  p.Complete()

  content, _ := json.MarshalIndent(p.ToTree(), "", "  ")
  return string(content)
}
