
package cirru

type charObj struct {
  x int
  y int
  text rune
  file *fileObj
}

func (char charObj) isBlank() bool {
  return char.text == ' '
}

func (char charObj) isOpenParen() bool {
  return char.text == '('
}

func (char charObj) isCloseParen() bool {
  return char.text == ')'
}

func (char charObj) isDollar() bool {
  return char.text == '$'
}

func (char charObj) isDoubleQuote() bool {
  return char.text == '"'
}

func (char charObj) isBackslash() bool {
  return char.text == '\\'
}
