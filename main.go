package main

import "fmt"

func main() {
	var stc Stack
	fmt.Print("capacity -> ")
	fmt.Println(stc.capacity)
	fmt.Println("\nVALUE | ELEMENT ADDR | PRIV ELEMENT ADDR")
	fmt.Println()

	for i := 0; i < 16; i++ {
		stc.push(i)
		//fmt.Println(stc.head)
		stc.printHead()
		fmt.Print("\n")
	}
	fmt.Print("stack head - > ")
	stc.printHead()
	fmt.Println()
	fmt.Print("capacity -> ")
	fmt.Println(stc.capacity)

	stc.pop()
	stc.pop()
	stc.pop()
	stc.pop()
	fmt.Println()

	fmt.Print("stack head (after stack.pop four times called...) - > ")
	stc.printHead()
	fmt.Println()
	fmt.Print("capacity -> ")
	fmt.Println(stc.capacity)
}