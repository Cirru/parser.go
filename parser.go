
package cirru

import (
  "strings"
  "fmt"
)

type fileObj struct {
  text string
  path string
}

func wrapText(text, filename string) (lines []inline) {
  for y, lineText := range strings.Split(text, "\n") {
    debugPrint("trace:", y, lineText)
    
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
  fmt.Println("")
  fmt.Println(xs...)
}