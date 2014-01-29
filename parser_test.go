
package cirru

import (
  "testing"
  "io/ioutil"
)

func TestParse(t *testing.T) {
  filename := "./cirru/grammar.cirru"
  code, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  codeText := string(code)
  res := Parse(codeText, filename)
  if err != nil {
    panic(err)
  }
  debugPrint(res)
}
