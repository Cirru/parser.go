
package cirru

func parseText(line inline, args List) List {
  tokens := tokenize(line.line)

  getBuffer := func (data tokenObj) Token {
    return data.buffer
  }

  var build func (byDollar bool) List
  build = func (byDollar bool) List {
    collection := List{}

    takeArgs := func () {
      if len(tokens) == 0 {
        if len(args) > 0 {
          collection = append(collection, args...)
          args = List{}
        }
      }
    }

    takeArgs()

    for {
      if len(tokens) == 0 {
        if byDollar {
          if len(tokens) > 0 && tokens[0].class == "closeParen" {
            return collection
          }
        }
        break
      }
      cursor := tokens[0]
      tokens = tokens[1:]
      switch cursor.class {
      case "string":
        collection = append(collection, getBuffer(cursor))
      case "text":
        if cursor.buffer.Text == "$" {
          collection = append(collection, build(true))
        } else {
          collection = append(collection, getBuffer(cursor))
        }
      case "openParen":
        collection = append(collection, build(false))
      case "closeParen":
        return collection
      }
      takeArgs()
    }
    return collection
  }
  return build(false)
}