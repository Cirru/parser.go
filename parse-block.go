
package cirru

func parseNested(currLines []inline) []interface{} {

  debugPrint("parsing nested:", currLines)

  for _, line := range currLines {
    line.dedent()
  }
  return parseBlock(currLines)
}

func parseBlock(currLines []inline) []interface{} {

  debugPrint("parsing block:", currLines)

  collection := []interface{}{}
  buffer := []inline{}

  digestBuffer := func () {
    if len(buffer) > 0 {
      line := buffer[0]
      var tree []interface{}
      if len(collection) == 0 && line.getIndent() > 0 {
        tree = parseNested(buffer)
      } else {
        tree = parseTree(buffer)
      }
      collection = append(collection, tree)
      buffer = []inline{}
    }
  }

  for _, line := range currLines {
    if line.isEmpty() {
      continue
    }
    if line.getIndent() == 0 {
      digestBuffer()
    }
    buffer = append(buffer, line)
  }
  digestBuffer()
  return collection
}

func parseTree(tree []inline) []interface{} {

  debugPrint("parsing tree", tree)

  for _, line := range tree[1:] {
    line.dedent()
  }
  args := []interface{}{}
  if len(tree[1:]) > 0 {
    args = parseBlock(tree[1:])
  }

  newTree := parseText(tree[0], args)
  return newTree
}