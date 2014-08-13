
package cirru

import "fmt"

type Parser struct {
  ast *[]*Expression
  state *state
}

func NewParser() Parser {
  list := &[]*Expression{}
  history := &[]interface{}{}
  emptyList := &[]interface{}{}
  first := &Expression{emptyList}
  *history = append(*history, first)
  *list = append(*list, first)
  mockToken := &Token{"", 1, 1, 1, 1}
  initial := &state{stateIndent, mockToken, 0, 1, 1, history, first}
  p := Parser{list, initial}
  return p
}

func (p *Parser) Read(c rune) {
  fmt.Printf("\n%v %s\n", p.state, string(c))
  if c == NewLine {
    p.state.countNewline()
  } else {
    p.state.countLetter()
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
  case stateString, stateEscape, stateToken: s.addBuffer(c)
  }
}

func (p *Parser) readParenLeft(c rune) {
  s := p.state
  switch s.name {
  case stateIndent:
    s.handleIndentation()
    s.pushStack()
  case stateString, stateEscape: s.addBuffer(c)
  case stateToken:
    s.completeToken()
    s.pushStack()
  }
}

func (p *Parser) readParenRight(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: panic("unexpected ParenRight at head")
  case stateString, stateEscape: s.addBuffer(c)
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
    s.completeSlash()
  case stateEscape, stateToken:
    s.addBuffer(c)
  }
}

func (p *Parser) GetAst() {
  fmt.Printf("%v\n\n", *(p.ast))
}

func (p *Parser) FormatAst() {
  for _, expr := range(*p.ast) {
    fmt.Printf("%v", *expr)
    println(expr.format())
  }
}

func (p *Parser) Complete() {
  p.state.completeToken()
}

func (p *Parser) ToJSON() (out []interface{}) {
  for _, child := range(*p.ast) {
    out = append(out, child.toJSON())
  }
  return
}