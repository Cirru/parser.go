
package cirru

func parseNested(currLines []inline) []interface{} {

  debugPrint("parsing nested:", currLines)

  newLines := []inline{}
  for _, line := range currLines {
    // this line is strange, array might be copied
    line.line = line.dedent()
    newLines = append(newLines, line)
  }
  return parseBlock(newLines)
}

func parseBlock(currLines []inline) []interface{} {

  debugPrint("parsing block:", currLines)

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
  debugPrint("block return", collection)
  return collection
}

func parseTree(tree []inline) []interface{} {

  debugPrint("parsing tree", tree)

  treeBlock := []inline{}
  for _, line := range tree[1:] {
    line.line = line.dedent()
    treeBlock = append(treeBlock, line)
  }
  args := []interface{}{}
  if len(treeBlock) > 0 {
    debugPrint("begin parseBlock:", treeBlock)
    args = parseBlock(treeBlock)
  }

  debugPrint("args to pass", args)
  newTree := parseText(tree[0], args)
  return newTree
}