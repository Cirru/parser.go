
Cirru Parser
------

Cirru Parser implemented in Go. Visit http://cirru.org for more.

You may also find one [writtern in CoffeeScript][parser].

[parser]: https://github.com/Cirru/cirru-parser.coffee

### Usage

[![GoDoc](https://godoc.org/github.com/Cirru/parser?status.png)](https://godoc.org/github.com/Cirru/parser)

You may find a complete demo at `parser_test.go`. Here is an overview.

```go
// import "github.com/Cirru/parser"

b, _ := ioutil.ReadFile("demo.cirru")

p := parser.NewParser()
for _, c := range b {
  p.Read(rune(c))
}
p.Complete()

content, _ := json.MarshalIndent(p.ToTree(), "", "  ")
string(content) // in JSON
```

### License

MIT