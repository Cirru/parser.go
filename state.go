
package cirru

type stateName int

const (
  stateIndent stateName = iota
  stateToken
  stateString
  stateEscape
)

type state struct {
  name stateName
  buffer Token
  level int
  x, y int
  history []interface{}
  cursor Expression
}

// state names
// indent token string escape

func (s *state) countNewline() {
  s.x = 0
  s.y += 1
}

func (s *state) countLetter() {
  s.x += 1
}

func (s *state) dropEmptyLine() {
  s.level = 0
  s.name = stateIndent
}

func (s *state) startBuffer() {
  buffer := s.buffer
  buffer.x = s.x
  buffer.ex = s.x
  buffer.y = s.y
  buffer.ey = s.y
  buffer.text = ""
}

func (s *state) completeBuffer() {
  buffer := s.buffer
  buffer.ex = s.x
  buffer.ey = s.y
  if len(buffer.text) > 0 {
    s.cursor.push(buffer)
  }
  s.startBuffer()
}

func (s *state) addBuffer(c rune) {
  buffer := s.buffer
  char := string(c)
  buffer.text += char
}

func (s *state) addIndentation() {
  s.level += 1
}

func (s *state) startString() {
  s.name = stateString
  s.buffer = Token{"", s.x, s.y, s.x, s.y}
}

func (s *state) completeString() {
  buffer := s.buffer
  buffer.ex = s.x
  buffer.ey = s.y
  if len(buffer.text) > 0 {
    s.cursor.push(buffer)
  }
  s.startString()
}

func (s *state) pushStack() {
  s.name = stateToken
  list := []interface{}{}
  expr := Expression{list}
  s.history = append(s.history, s.cursor)
  s.cursor = expr
}

func (s *state) popStack() {
  s.name = stateToken
  endIndex := len(s.history) - 1
  if expr, ok := s.history[endIndex].(Expression); ok {
    s.cursor = expr
  } else {
    panic("got wrong thing from history")
  }
  s.history = s.history[:endIndex]
}

func (s *state) handleIndentation() {
  indented := len(s.buffer.text)
  if indented > s.level {
    diff := indented - s.level
    for ; diff > 0; diff -= 2 {
      s.pushStack()
      if diff == 1 {
        panic("odd indentation not valid")
      }
    }
  } else if indented < s.level {
    diff := s.level - indented
    for ; diff > 0; diff -= 2 {
      s.popStack()
      if diff == 1 {
        panic("odd indentation not valid")
      }
    }
  }
  s.level = indented
}
