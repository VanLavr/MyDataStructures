package main

import "fmt"

func main() {
	
	// stack check
	var stc Stack
	fmt.Print("capacity -> ")
	fmt.Println(stc.capacity)
	fmt.Println("\nVALUE | ELEMENT ADDR | PRIV ELEMENT ADDR")
	fmt.Println()

	for i := 0; i < 16; i++ {
		stc.push(i)
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


	// dequeue check
	var dq Dequeue
	fmt.Print("\n\ncapacity -> ")
	fmt.Println(dq.capacity)
	fmt.Println("\nVALUE | ELEMENT ADDR | PRIV ELEMENT ADDR | NEXT ELEMENT ADDR")
	fmt.Println()

	for i := 0; i < 16; i++ {
		if i % 2 == 0 {
			dq.PushUp(i)
			dq.printHead()
			fmt.Println()
		} else {
			dq.PushDown(i)
			dq.printTail()
			fmt.Println()
		}
	}
	fmt.Print("dequeue head - > ")
	dq.printHead()
	fmt.Println()
	fmt.Print("dequeue tail - > ")
	dq.printTail()
	fmt.Println()
	fmt.Print("capacity -> ")
	fmt.Println(dq.capacity)

	dq.PopUp()
	dq.PopDown()
	dq.PopUp()
	dq.PopDown()
	fmt.Println()

	fmt.Print("dequeue head (after dequeue.PopUp two times and dequeue.PopDown two times called...) - > ")
	dq.printHead()
	fmt.Println()
	fmt.Print("dequeue tail (after dequeue.PopUp two times and dequeue.PopDown two times called...) - > ")
	dq.printTail()
	fmt.Println()
	fmt.Print("capacity -> ")
	fmt.Println(dq.capacity)
	fmt.Println("how actually dequeue is look like before deleting some elements:\nhead --> 14 12 10 8 6 4 2 0 1 3 5 7 9 11 13 15 <-- tail\nand how it looks like after deleting some elements:\nhead --> 10 8 6 4 2 0 1 3 5 7 9 11 <-- tail")
}