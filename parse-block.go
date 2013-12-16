
package cirru

func parseNested(currLines []inline) []interface{} {

  newLines := []inline{}
  for _, line := range currLines {
    line.line = line.dedent()
    newLines = append(newLines, line)
  }
  return parseBlock(newLines)
}

func parseBlock(currLines []inline) []interface{} {

  collection := []interface{}{}
  lines := []inline{}

  digestBuffer := func () {
    if len(lines) > 0 {
      line := lines[0]
      var tree []interface{}
      if len(collection) == 0 && line.getIndent() > 0 {
        tree = parseNested(lines)
      } else {
        tree = parseTree(lines)
      }
      collection = append(collection, tree)
      lines = []inline{}
    }
  }

  for _, line := range currLines {
    if line.isEmpty() {
      continue
    }
    if line.getIndent() == 0 {
      digestBuffer()
    }
    lines = append(lines, line)
  }
  digestBuffer()
  return collection
}

func parseTree(tree []inline) []interface{} {

  treeBlock := []inline{}
  for _, line := range tree[1:] {
    line.line = line.dedent()
    treeBlock = append(treeBlock, line)
  }
  args := []interface{}{}
  if len(treeBlock) > 0 {
    args = parseBlock(treeBlock)
  }

  newTree := parseText(tree[0], args)
  return newTree
}