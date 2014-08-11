
package cirru

import "fmt"

type Parser struct {
  Done bool
  AST []Expression
  state state
}

func NewParser() Parser {
  list := []Expression{}
  history := []interface{}{}
  first := Expression{}
  list = append(list, first)
  mockToken := Token{"", 0, 0, 0, 0}
  initial := state{stateIndent, mockToken, 0, 0, 0, history, first}
  p := Parser{false, list, initial}
  return p
}

func (p *Parser) Read(c rune) {
  // fmt.Printf("\n%v", p.state)
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
  case stateToken: s.completeBuffer()
  }
}

func (p *Parser) readSpace(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: s.addIndentation()
  case stateString: p.readCode(c)
  case stateEscape: panic("no need to use Space in escape")
  case stateToken: s.completeBuffer()
  }
}

func (p *Parser) readCode(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: s.handleIndentation()
  case stateString: s.addBuffer(c)
  case stateEscape: s.addBuffer(c)
  case stateToken: s.addBuffer(c)
  }
}

func (p *Parser) readParenLeft(c rune) {
  s := p.state
  switch s.name {
  case stateIndent:
    s.handleIndentation()
    s.pushStack()
  case stateString: s.addBuffer(c)
  case stateEscape: s.addBuffer(c)
  case stateToken:
    s.completeBuffer()
    s.pushStack()
  }
}

func (p *Parser) readParenRight(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: panic("unexpected ParenRight at head")
  case stateString: s.addBuffer(c)
  case stateEscape: s.addBuffer(c)
  case stateToken:
    s.completeBuffer()
    s.popStack()
  }
}

func (p *Parser) readQuote(c rune) {
  println("read Quote")
}

func (p *Parser) readBackslash(c rune) {
  println("read Backslash")
}

func (p *Parser) GetAst() {
  fmt.Printf("%v\n\n", p.AST)
}