
package cirru

type Expression struct {
  list *[]interface{}
}

func (e *Expression) push(child interface{}) {
  *e.list = append(*e.list, child)
}

func (e *Expression) toArray() (out []interface{}) {
  out = []interface{}{}
  for _, child := range(*e.list) {
    if expr, ok := child.(*Expression); ok {
      out = append(out, expr.toArray())
    } else if token, ok := child.(Token); ok {
      out = append(out, token)
    }
  }
  return
}

func (e *Expression) toTree() (out []interface{}) {
  out = []interface{}{}
  for _, child := range(*e.list) {
    if expr, ok := child.(*Expression); ok {
      out = append(out, expr.toTree())
    } else if token, ok := child.(Token); ok {
      out = append(out, token.getText())
    }
  }
  return
}

func (e *Expression) resolveDollar() {
  for i, child := range(*e.list) {
    if expr, ok := child.(*Expression); ok {
      expr.resolveDollar()
    } else if token, ok := child.(Token); ok {
      if token.getText() == string(Dollar) {
        former := (*e.list)[(i+1):]
        childExpr := &Expression{&former}
        childExpr.resolveDollar()
        *e.list = append((*e.list)[:i], childExpr)
        break
      }
    }
  }
}

func (e *Expression) resolveComma() {
  oldList := e.list
  e.list = &[]interface{}{}
  for i, child := range(*oldList) {
    if expr, ok := child.(*Expression); ok {
      if i == 0 {
        expr.resolveComma()
        e.push(expr)
        continue
      }
      if expr.hasLeadingComma() {
        expr.resolveComma()
        for j, item := range(*expr.list) {
          if j == 0 {
            continue
          }
          if childExpr, ok := item.(*Expression); ok {
            childExpr.resolveComma()
            e.push(childExpr)
          } else {
            e.push(item)
          }
        }
      } else {
        expr.resolveComma()
        e.push(expr)
      }
    } else {
      e.push(child)
    }
  }
}

func (e *Expression) hasLeadingComma() bool {
  if len(*e.list) == 0 {
    return false
  }
  if token, ok := (*e.list)[0].(Token); ok {
    return token.getText() == string(Comma)
  } else {
    return false
  }
}
