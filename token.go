
package cirru

type Token struct {
  text string
  x, y, ex, ey int
}

func (t *Token) empty() {
  t.text = ""
}

func (t *Token) format() string {
  return t.text
}

func (t *Token) add(c rune) {
  str := string(c)
  t.text += str
}

func (t *Token) toJSON() string {
  return t.text
}