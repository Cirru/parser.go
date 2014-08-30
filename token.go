
package parser

type Token struct {
  Text string `json:"text"`
  X int `json:"x"`
  Y int `json:"y"`
  Ex int `json:"ex"`
  Ey int `json:"ey"`
}

func (t *Token) empty() {
  t.Text = ""
}

func (t *Token) add(c rune) {
  str := string(c)
  t.Text += str
}

func (t *Token) getText() string {
  return t.Text
}

func (t *Token) setXy(x, y int) {
  t.X = x
  t.Y = y
}

func (t *Token) setExy(x, y int) {
  t.Ex = x
  t.Ey = y
}

func (t *Token) len() int {
  return len(t.Text)
}