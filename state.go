
package parser

type state struct {
  name stateName
  buffer *Token
  level int
  x, y int
  history *[]*Expression
  cursor *Expression
  file string
}

// state names
// indent token string escape

func (s *state) countNewline() {
  s.x = 1
  s.y += 1
}

func (s *state) countLetter() {
  s.x += 1
}

func (s *state) dropEmptyLine() {
  s.buffer.empty()
}

func (s *state) beginToken() {
  s.name = stateToken
  buffer := s.buffer
  buffer.setXy(s.x, s.y)
  buffer.setExy(s.x, s.y)
  buffer.empty()
}

func (s *state) completeToken() {
  buffer := s.buffer
  buffer.setXy(s.x, s.y)
  if buffer.len() > 0 {
    s.cursor.push(*buffer)
  }
  s.beginToken()
}

func (s *state) addBuffer(c rune) {
  s.buffer.add(c)
}

func (s *state) beginString() {
  s.name = stateString
  s.buffer = &Token{"", s.x, s.y, s.x, s.y, s.file}
}

func (s *state) completeString() {
  buffer := s.buffer
  buffer.setExy(s.y, s.y)
  if len(buffer.getText()) > 0 {
    s.cursor.push(*buffer)
  }
  s.beginToken()
}

func (s *state) pushStack() {
  s.name = stateToken
  list := &[]interface{}{}
  expr := &Expression{list}
  *s.history = append(*s.history, s.cursor)
  s.cursor.push(expr)
  s.cursor = expr
}

func (s *state) popStack() {
  if len(*s.history) == 0 {
    return
  }
  endIndex := len(*s.history) - 1
  expr := (*s.history)[endIndex]
  s.name = stateToken
  s.cursor = expr
  *s.history = (*s.history)[:endIndex]
}

func (s *state) handleIndentation() {
  indented := s.buffer.len()
  if indented > s.level {
    diff := indented - s.level
    if diff % 2 == 1 {
        panic("odd indentation not valid")
    }
    for ; diff > 0; diff -= 2 {
      s.pushStack()
    }
  } else if indented < s.level {
    diff := s.level - indented
    if diff % 2 == 1 {
        panic("odd indentation not valid")
    }
    for ; diff >= 0; diff -= 2 {
      s.popStack()
    }
    s.pushStack()
  } else {
    s.popStack()
    s.pushStack()
  }
  s.level = indented
}

func (s *state) beginEscape() {
  s.name = stateEscape
}

func (s *state) beginNewline() {
  s.name = stateIndent
  buffer := s.buffer
  buffer.setXy(s.x, s.y)
  buffer.setExy(s.x, s.y)
}

func (s *state) completeEscape() {
  s.name = stateString
}

func (s *state) getName() string {
  switch s.name {
  case 0: return "Indent"
  case 1: return "Token"
  case 2: return "String"
  case 3: return "Escape"
  }
  return "<unknown>"
}