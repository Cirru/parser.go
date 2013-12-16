
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

func wrapText(text, filename string) (lines []inline) {
  for y, lineText := range strings.Split(text, "\n") {
    
    charList := []charObj{}
    line := inline{charList}
    file := &fileObj{text, filename}
    lineText = strings.TrimRight(lineText, " ")
    for x, charText := range strings.Split(lineText, "") {
      char := charObj{x, y, rune(charText[0]), file}
      line.line = append(line.line, char)
    }
    lines = append(lines, line)
  }
  return lines
}

func Parse(code, filename string) []interface{} {
  lines := wrapText(code, filename)
  return parseBlock(lines)
}

func debugPrint(xs ...interface{}) {
  list := []interface{}{}
  for _, item := range xs {
    json, err := json.MarshalIndent(item, "", "  ")
    if err != nil {
      panic(err)
    }
    list = append(list, interface{}(string(json)))
  }
  fmt.Println("")
  fmt.Println("")
  fmt.Println(xs...)
  fmt.Println(list...)
}
