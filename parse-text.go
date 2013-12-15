
package cirru

func parseText(line inline, args []interface{}) []interface{} {
  debugPrint("start parsing text", line, args)
  tokens := tokenize(line.line)

  getBuffer := func (data tokenObj) bufferObj {
    return data.buffer
  }

  var build func (byDollar bool) []interface{}
  build = func (byDollar bool) []interface{} {
    collection := []interface{}{}

    takeArgs := func () {
      if len(tokens) == 0 {
        if len(args) > 0 {
          debugPrint("args:", args[0])
          collection = append(collection, args...)
          args = []interface{}{}
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
        if cursor.buffer.text == "$" {
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