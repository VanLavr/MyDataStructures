package main

import "fmt"

type Dequeue struct {
	capacity int
	head *ElementOfDequeue
	tail *ElementOfDequeue
}

func (d *Dequeue) PushUp(element int) int {
	if d.capacity == 0 {

		var el ElementOfDequeue
		el.value = element
		el.pointerOnPtiviousElement = nil
		el.pointerOnNextElement = nil
		el.pointerOnThisElement = &el

		d.capacity++
		d.head = &el
		d.tail = &el

		return d.capacity

	} else {

		var el ElementOfDequeue
		el.value = element
		el.pointerOnPtiviousElement = d.head
		el.pointerOnNextElement = nil
		el.pointerOnThisElement = &el

		d.capacity++
		d.head = &el
		
		return d.capacity

	}
}

func (d *Dequeue) PushDown(element int) int {
	if d.capacity == 0 {

		var el ElementOfDequeue
		el.value = element
		el.pointerOnPtiviousElement = nil
		el.pointerOnNextElement = nil
		el.pointerOnThisElement = &el

		d.capacity++
		d.head = &el
		d.tail = &el

		return d.capacity

	} else {

		var el ElementOfDequeue
		el.value = element
		el.pointerOnPtiviousElement = nil
		el.pointerOnNextElement = d.tail
		el.pointerOnThisElement = &el

		d.capacity++
		d.tail = &el
		
		return d.capacity

	}
}

func (d *Dequeue) PopUp() (capacity int, err error) {
	if d.capacity == 0 {
		return 0, err
	} else {
		d.head = d.head.pointerOnPtiviousElement
		d.capacity--
		return d.capacity, nil
	}
}

func (d *Dequeue) PopDown() (capacity int, err error) {
	if d.capacity == 0 {
		return 0, err
	} else {
		d.tail = d.tail.pointerOnNextElement
		d.capacity--
		return d.capacity, nil
	}
}

func (d *Dequeue) printHead() {
	fmt.Print(*d.head)
}

func (d *Dequeue) printTail() {
	fmt.Print(*d.tail)
}