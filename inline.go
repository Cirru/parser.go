
package cirru

import (
)

type inline struct {
  line []charObj
}

func (line inline) isEmpty() bool {
  if len(line.line) == 0 {
    return true
  } else {
    return false
  }
}

func (line inline) getIndent() int {
  n := 0
  for _, char := range(line.line) {
    if char.isBlank() {
      n += 1
    } else {
      break
    }
  }
  return (n + 1) / 2
}

func (line inline) dedent() []charObj {
  line.line = line.line[1:]
  if len(line.line) > 0 {
    first := line.line[0]
    if first.isBlank() {
      line.line = line.line[1:]
    }
  }
  return line.line
}

func (line inline) lineEnd() bool {
  return len(line.line) == 0
}