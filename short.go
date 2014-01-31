
package cirru

// ParseShort rids result from Parse but without file infos.
func ParseShort(code, filename string) (ret List) {
  result := Parse(code, filename)
  ret = short(result)
  return
}

func short(tree List) (ret List) {
  for _, item := range tree {
    if childList, ok := item.(List); ok {
      ret = append(ret, short(childList))
    } else if token, ok := item.(Token); ok {
      ret = append(ret, token.Text)
    }
  }
  return
}