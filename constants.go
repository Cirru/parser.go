
package parser

const newLine = rune('\n')
const space = rune(' ')

const parenLeft = rune('(')
const parenRight = rune(')')

const quote = rune('"')
const backslash = rune('\\')

const comma = rune(',')
const dollar = rune('$')

type stateName int

const (
  stateIndent stateName = iota
  stateToken
  stateString
  stateEscape
)
