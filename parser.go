
package cirru

// import "fmt"

type Parser struct {
  ast *Expression
  state *state
}

// Creates new parser which contains the following methods.
func NewParser() Parser {
  emptyList := &[]interface{}{}
  firstExpr := &Expression{emptyList}

  history := &[]*Expression{}

  mockToken := &Token{"", 1, 1, 1, 1}
  initialState := &state{stateIndent,
    mockToken,
    0, 1, 1,
    history,
    firstExpr,
  }
  p := Parser{firstExpr, initialState}
  return p
}

// Takes each character in rune type, and triggers parser.
func (p *Parser) Read(c rune) {
  s := p.state
  // safeChar := fmt.Sprintf("%q", c)
  // safeBuffer := fmt.Sprintf("%q", string(s.buffer.text))
  // println(s.getName(), "\t", safeChar, "\t", safeBuffer)
  if c == NewLine {
    s.countNewline()
  } else {
    s.countLetter()
  }
  switch c {
  case NewLine: p.readNewline(c)
  case Space: p.readSpace(c)
  case ParenLeft: p.readParenLeft(c)
  case ParenRight: p.readParenRight(c)
  case Quote: p.readQuote(c)
  case Backslash: p.readBackslash(c)
  default: p.readCode(c)
  }
}

func (p *Parser) readNewline(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: s.dropEmptyLine()
  case stateString: panic("unexpected NewLine in string")
  case stateEscape: panic("unexpected NewLine in escape")
  case stateToken:
    s.completeToken()
    s.beginNewline()
  }
}

func (p *Parser) readSpace(c rune) {
  s := p.state
  switch s.name {
  case stateIndent, stateString: s.addBuffer(c)
  case stateEscape: panic("no need to use Space in escape")
  case stateToken: s.completeToken()
  }
}

func (p *Parser) readCode(c rune) {
  s := p.state
  switch s.name {
  case stateIndent:
    s.handleIndentation()
    s.beginToken()
    s.addBuffer(c)
  case stateString, stateToken: s.addBuffer(c)
  case stateEscape:
    s.addBuffer(c)
    s.completeEscape()
  }
}

func (p *Parser) readParenLeft(c rune) {
  s := p.state
  switch s.name {
  case stateIndent:
    s.handleIndentation()
    s.pushStack()
    s.beginToken()
  case stateString: s.addBuffer(c)
  case stateEscape:
    s.addBuffer(c)
    s.completeEscape()
  case stateToken:
    s.completeToken()
    s.pushStack()
  }
}

func (p *Parser) readParenRight(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: panic("unexpected ParenRight at head")
  case stateString: s.addBuffer(c)
  case stateEscape:
    s.addBuffer(c)
    s.completeEscape()
  case stateToken:
    s.completeToken()
    s.popStack()
  }
}

func (p *Parser) readQuote(c rune) {
  s := p.state
  switch s.name {
  case stateIndent:
    s.handleIndentation()
    s.beginString()
  case stateString:
    s.completeString()
  case stateEscape:
    s.addBuffer(c)
    s.completeEscape()
  case stateToken:
    s.completeToken()
    s.beginString()
  }
}

func (p *Parser) readBackslash(c rune) {
  s := p.state
  switch s.name {
  case stateIndent:
    s.handleIndentation()
    s.beginToken()
  case stateString:
    s.beginEscape()
  case stateEscape, stateToken:
    s.addBuffer(c)
    s.completeEscape()
  }
}

// Tell the parser it's completed.
// Call this method when files reach the end.
func (p *Parser) Complete() {
  p.state.completeToken()
  p.ast.resolveDollar()
  p.ast.resolveComma()
}

// Get array out of a parser.
// It actually returns slices, but easy to mashaled into JSON.
func (p *Parser) ToArray() (out []interface{}) {
  return p.ast.toArray()
}