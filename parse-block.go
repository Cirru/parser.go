
package cirru

func parseNested(currLines []inline) List {

  newLines := []inline{}
  for _, line := range currLines {
    line.line = line.outdent()
    newLines = append(newLines, line)
  }
  return parseBlock(newLines)
}

func parseBlock(currLines []inline) List {

  collection := List{}
  lines := []inline{}
  count := 0
  empty := List{}
  track := make(chan bool)

  digestBuffer := func () {
    if len(lines) > 0 {
      collection = append(collection, empty)

      go func(lines []inline, length int) {
        line := lines[0]
        var tree List
        if (length - 1) == 0 && line.getIndent() > 0 {
          tree = parseNested(lines)
        } else {
          tree = parseTree(lines)
        }
        collection[length - 1] = tree
        track <- true
      }(lines, len(collection))

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

  for {
    <- track
    count += 1
    if count == len(collection) {
      break
    }
  }

  return collection
}

func parseTree(tree []inline) List {

  treeBlock := []inline{}
  for _, line := range tree[1:] {
    line.line = line.outdent()
    treeBlock = append(treeBlock, line)
  }
  args := List{}
  if len(treeBlock) > 0 {
    args = parseBlock(treeBlock)
  }

  newTree := parseText(tree[0], args)
  return newTree
}