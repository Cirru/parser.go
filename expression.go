
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
    if expr, ok := child.(Expression); ok {
      println("found expr")
      out += expr.format()
    } else if token, ok := child.(Token); ok {
      println("found token")
      out += "'"
      out += token.format()
      out += "' "
    }
  }
  out += ")"
  return
}