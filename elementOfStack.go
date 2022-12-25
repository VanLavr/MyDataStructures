package main

type ElementOfStack struct {
	value int // value of current element
	pointerOnThisElement *ElementOfStack // this field of elemnet contains pointer on itself (self adderess)
	pointerOnPtiviousElement *ElementOfStack // this field points on privious element in stack
}

// comment