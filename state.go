
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

func (s *state) handleIndentation() {
}

func (s *state) indent(n int) {
  if n <= 0 {
    panic("n <= 0")
  }
}

func (s *state) unindent(n int) {
  if n >= 0 {
    panic("n >= 0")
  }
}
