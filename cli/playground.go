
package main

import "fmt"

type a struct {
  list *[]int
}

func main() {
  mockList := &[]int{1,2,3}
  nameA := &a{mockList}
  nameB := nameA
  *nameB.list = append(*(nameB.list), 4)
  fmt.Printf("%v", *nameA.list)
}