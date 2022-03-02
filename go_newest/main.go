package main

import "fmt"

func print[T any] (arr []T){
  for _, v := range arr {
    fmt.Print(v)
    fmt.Print(" ")
  }
  fmt.Println("")
}
func simpleTest(){
  strs := []string{"HelloWorld:","Guys!"}
  decs := []float64{3.14, 1.14, 1.618, 2.718 }
  nums := []int{2,4,6,8}

  print(strs)
  print(decs)
  print(nums)
}

type stack [T any] []T
func (s *stack[T]) push(elem T) {
  *s = append(*s, elem)
}
func (s *stack[T]) pop() {
  if len(*s) > 0 {
    *s = (*s)[:len(*s)-1]
  } 
}
func (s *stack[T]) top() *T{
  if len(*s) > 0 {
    return &(*s)[len(*s)-1]
  } 
  return nil
}
func (s *stack[T]) len() int{
  return len(*s)
}
func (s *stack[T]) print() {
  for _, elem := range *s {
    fmt.Print(elem)
    fmt.Print(" ")
  }
  fmt.Println("")
}
func multiTest(){
  ss := stack[string]{}
  ss.push("Hello")
  ss.push("Hao")
  ss.push("Chen")
  ss.print()
  fmt.Printf("stack top is - %v\n", *(ss.top()))
  ss.pop()
  ss.pop()
  ss.print()
  
  ns := stack[int]{}
  ns.push(10)
  ns.push(20)
  ns.print()
  ns.pop()
  ns.print()
  *ns.top() += 1
  ns.print()
  ns.pop()
  fmt.Printf("stack top is - %v\n", ns.top())
}

// 编译执行:
// go run -gcflags=-G=3 ./main.go
func main() {
  simpleTest()
  multiTest()
}