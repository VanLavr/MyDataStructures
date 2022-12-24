package main

import "fmt"

func main() {
	var stc Stack

	for i := 0; i < 10; i++ {
		stc.push(i)
		//fmt.Println(stc.head)
		stc.printHead()
		fmt.Print(" ")
	}

	stc.pop()
	fmt.Println()

	for i := 0; i < 9; i++ {
		stc.push(i)
		//fmt.Println(stc.head)
		stc.printHead()
		fmt.Print(" ")
	}
}