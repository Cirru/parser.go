
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
  case NewLine: p.readNewline()
  case Space: p.readSpace(c)
  case ParenLeft: p.readParenLeft()
  case ParenRight: p.readParenRight()
  case Quote: p.readQuote()
  case Backslash: p.readBackslash()
  default: p.readCode(c)
  }
}

func (p *Parser) readNewline() {
  s := p.state
  switch s.name {
  case stateIndent: s.dropEmptyLine()
  case stateString: panic("unexpacted NewLine in string")
  case stateEscape: panic("unexpected NewLine in escape")
  case stateToken: s.completeBuffer()
  }
}

func (p *Parser) readSpace(c rune) {
  s := p.state
  switch s.name {
  case stateIndent: p.state.addIndentation()
  case stateString: p.readCode(c)
  case stateEscape: panic("no need to use Space in escape")
  case stateToken: p.state.completeBuffer()
  }
}

func (p *Parser) readCode(c rune) {
  println("normal char")
}

func (p *Parser) readParenLeft() {
  println("ParenLeft")
}

func (p *Parser) readParenRight() {
  println("ParenRight")
}

func (p *Parser) readQuote() {
  println("read Quote")
}

func (p *Parser) readBackslash() {
  println("read Backslash")
}

func (p *Parser) GetAst() {
  fmt.Printf("%v\n\n", p.AST)
}