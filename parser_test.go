
package cirru

import (
  "testing"
  "io/ioutil"
  "encoding/json"
  "strings"
)


func TestParse(t *testing.T) {
  names := []string {
    "demo",
    "folding",
    "indent",
    "line",
    "parentheses",
    "quote",
    "unfolding",
  }
  for _, filename := range names {
    codename := "./cirru/" + filename + ".cirru"
    jsonname := "./cirru/" + filename + ".json"
    codeContent, _ := ioutil.ReadFile(codename)
    jsonContent, _ := ioutil.ReadFile(jsonname)

    codeString := string(codeContent)

    jsonString := strings.TrimFunc(string(jsonContent), func(x rune) bool {
      if string(x) == " " {
        return true
      }
      if string(x) == "\n" {
        return true
      }
      return false
    })

    tree := ParseShort(codeString, codename)
    result, _ := json.MarshalIndent(tree, "", "  ")
    resultString := string(result)
    resultString = strings.Replace(resultString, "\\u003c", "<", -1)
    resultString = strings.Replace(resultString, "\\u003e", ">", -1)
    if jsonString != resultString {
      println("not equal", filename)
      println(jsonString)
      println(resultString)
    }
  }
}
