package main

import (
	"encoding/json"
	"fmt"
)

type HI struct {
	Error bool
	Name  string
}

func main() {
	h := new(HI)
	h.Error = false
	h.Name = "Tom"

	b, err := json.Marshal(h)
	fmt.Println(err)
	fmt.Println(b)

	g := new(HI)
	err = json.Unmarshal(b, g)
	fmt.Println(err)
	fmt.Println(g)
}
