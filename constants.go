
package cirru

const NewLine = rune('\n')
const Space = rune(' ')

const ParenLeft = rune('(')
const ParenRight = rune(')')

const Quote = rune('"')
const Backslash = rune('\\')

const Comma = rune(',')
const Dollar = rune('$')

type stateName int

const (
  stateIndent stateName = iota
  stateToken
  stateString
  stateEscape
)
