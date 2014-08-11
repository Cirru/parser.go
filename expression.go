
package cirru

type Expression struct {
  list *[]interface{}
}

func (e *Expression) push(child interface{}) {
  *e.list = append(*e.list, child)
}