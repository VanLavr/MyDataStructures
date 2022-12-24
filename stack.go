package main

import "fmt"

type Stack struct {
	capacity int
	head *Element
	firstElem Element
}

func (s *Stack) push(element int) int {
	if s.capacity == 0 {
		
		var el Element
		el.value = element
		el.pointerOnThisElement = &el
		el.pointerOnPtiviousElement = nil

		s.capacity++
		s.head = &el
		s.firstElem.value = element

		return s.capacity
	
	} else {

		var el Element
		el.value = element
		el.pointerOnThisElement = &el
		el.pointerOnPtiviousElement = s.head

		s.capacity++
		s.head = &el
		s.firstElem.value = element

		return s.capacity
	}
}

func (s *Stack) pop() (capacity int, err error) {
	if s.capacity == 0 {
		return 0, err
	} else {
		s.head = nil
		s.capacity--
		return s.capacity, nil
	}
}

func (s *Stack) printHead() {
	fmt.Print(s.firstElem.value)
}