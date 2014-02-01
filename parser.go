
// Package cirru parses code in Cirru Grammer into a tree.
// That tree could be later interpreted in the runtime.
// Cirru is designed for making small scripting tools.
// Currently, this code is in developing mode.
package cirru

import (
  "strings"
  "fmt"
  "encoding/json"
)

type fileObj struct {
  text string
  path string
}

// List is an abstraction on []interface{} of cirru
// List consists of Token and List
type List []interface{}

func wrapText(text, filename string) (lines []inline) {
  for y, lineText := range strings.Split(text, "\n") {
    
    charList := []charObj{}
    line := inline{charList}
    file := &fileObj{text, filename}
    lineText = strings.TrimRight(lineText, " ")
    if len(lineText) == 0 {
      continue
    }
    for x, charText := range strings.Split(lineText, "") {
      char := charObj{x, y, rune(charText[0]), file}
      line.line = append(line.line, char)
    }
    lines = append(lines, line)
  }
  return lines
}

// Parse returns value is a slice mixed with strings and slices
func Parse(code, filename string) (ret List) {
  lines := wrapText(code, filename)
  if len(lines) != 0 {
    ret = parseBlock(lines)
  }
  return
}

func debugPrint(xs ...interface{}) {
  list := List{}
  for _, item := range xs {
    jsonContent, err := json.MarshalIndent(item, "", "  ")
    if err != nil {
      panic(err)
    }
    list = append(list, interface{}(string(jsonContent)))
  }
  fmt.Println("")
  fmt.Println("")
  fmt.Println(xs...)
  fmt.Println(list...)
}
