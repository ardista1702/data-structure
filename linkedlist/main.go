package main

import (
	"fmt"
	"linkedlist/listnode"
)

func main() {
	list := listnode.NewList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Prepend("x")
	list.Prepend("y")
	list.Prepend("z")
	arr := list.ToArray()
	fmt.Println(arr)
	fmt.Println(list.Size())
	fmt.Println(len(arr) == list.Size())
	list.Delete("a")
	fmt.Println(list.ToArray())
	fmt.Println(list.Size())
}
