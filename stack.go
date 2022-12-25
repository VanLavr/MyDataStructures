/*
FOR KOLYA
interprite this structure and functions like kinda class with two methods
waiting for u)
*/

package main

import "fmt"

type Stack struct {
	capacity int
	head *ElementOfStack // points on top of the stack (upper element)
}

func (s *Stack) push(element int) int { // push method...
	if s.capacity == 0 { // checking capacity itself and if its equal zero, pointer on the privious element for our first element will be equal nil (null)
		                                                                                           //       /*\ 
		var el ElementOfStack // creating element                                                            |
		el.value = element // assing our element value as it was given as argument in method                 |
		el.pointerOnThisElement = &el // assign pointer on this element with self address                    |
		el.pointerOnPtiviousElement = nil // assigned nil----------------------------------------------------!

		s.capacity++ // extending capacity...
		s.head = &el // this element now is the head of our stack

		return s.capacity
	
	} else { // here we did all the same stuff, but assigned pointer on ptivious element with address of privious element

		var el ElementOfStack
		el.value = element
		el.pointerOnThisElement = &el
		el.pointerOnPtiviousElement = s.head

		s.capacity++
		s.head = &el

		return s.capacity
	}
}

func (s *Stack) pop() (capacity int, err error) { // pop method...
	if s.capacity == 0 { // returning error if there are no element in stack
		return 0, err
	} else { // here (if you are writing it in C++ you should delete object, but in Go i can just assign pointer on my object with nil (null) and garbage collector will delete it)
		s.head = s.head.pointerOnPtiviousElement
		s.capacity--
		return s.capacity, nil
	}
}

func (s *Stack) printHead() { // just printing element...
	fmt.Print(*s.head)
}