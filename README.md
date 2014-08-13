
Cirru Parser
------

Cirru Parser implemented in Go. Visit http://cirru.org for more.

You may also find one [writtern in CoffeeScript][parser].

[parser]: https://github.com/Cirru/cirru-parser.coffee

### Usage

[![GoDoc](https://godoc.org/github.com/Cirru/cirru-parser.go?status.png)](https://godoc.org/github.com/Cirru/cirru-parser.go)

You may find a complete demo at `cli/main.go`. Here is an overview.

```go
import "github.com/Cirru/cirru-parser.go"

b, _ := ioutil.ReadFile("demo.cirru")

parser := cirru.NewParser()
for _, c := range b {
  parser.Read(rune(c))
}
parser.Complete()

content, _ := json.MarshalIndent(parser.ToArray(), "", "  ")
gotAst := string(content)
```

### License

MIT