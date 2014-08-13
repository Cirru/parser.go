
package cirru

type Expression struct {
  list *[]interface{}
}

func (e *Expression) push(child interface{}) {
  *e.list = append(*e.list, child)
}

func (e *Expression) insert(child *Expression) {
  *e.list = append(*e.list, child)
}

func (e *Expression) format() (out string) {
  out += "("
  for _, child := range(*e.list) {
    if expr, ok := child.(*Expression); ok {
      out += expr.format()
    } else if token, ok := child.(Token); ok {
      out += "'"
      out += token.format()
      out += "' "
    }
  }
  out += ")"
  return
}

func (e *Expression) toJSON() (out []interface{}) {
  for _, child := range(*e.list) {
    if expr, ok := child.(*Expression); ok {
      out = append(out, expr.toJSON())
    } else if token, ok := child.(Token); ok {
      out = append(out, token.toJSON())
    } else {
      panic("unexpected type got from AST")
    }
  }
  return
}