package main

import (
	"fmt"
	"lesson09/pkg/binary"
)

func main() {
	bin := binary.New()
	el := binary.Element{
		Value: 1,
		ID:    3,
	}
	bin.Insert(&el)
	el2 := binary.Element{
		Value: 1,
		ID:    1,
	}
	bin.Insert(&el2)
	r, _ := bin.Search(1)
	fmt.Printf("%+v\n", r)
	r2, _ := bin.Search(5)
	fmt.Printf("%+v\n", r2)
}
