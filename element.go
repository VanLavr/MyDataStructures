package main

type Element struct {
	value int // value of current element
	pointerOnThisElement *Element // this field of elemnet contains pointer on itself (self adderess)
	pointerOnPtiviousElement *Element // this field points on privious element in stack
}

// changes in new branch...