package main

import(
"fmt"
)

type Hello interface {
	Hi(x int) int
	Say(y string) string
}

type A struct {
	ID int
	Name string
}

func (a A)Hi(x int) int {
	return a.ID+x
}

func (a A)Say(y string) string {
	return a.Name+y
}

func main(){
var a1 =A{
	ID:1,
	Name:"Tom",
}
var h Hello
h=a1
x:=h.Hi(1)
y:=h.Say("er")
fmt.Println(x,y)
}

