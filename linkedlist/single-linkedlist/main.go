package main

import (
	"fmt"
	"single-linkedlist/listnode"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	list_ := listnode.NewList()
	for _, num := range nums {
		list_.Insert(num)
	}
	current := list_.Head

	for current != nil {
		fmt.Println("node:", &current)
		fmt.Printf("current address: %p\n", current)
		fmt.Println("current value:", current.Val)
		fmt.Printf("next address: %p\n", current.Next)
		fmt.Println("==========")
		current = current.Next
	}
	
}
