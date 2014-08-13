
package cirru

type Token struct {
  text string
  x, y, ex, ey int
}

func (t *Token) empty() {
  t.text = ""
}

func (t *Token) add(c rune) {
  str := string(c)
  t.text += str
}

func (t *Token) getText() string {
  return t.text
}