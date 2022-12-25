package main

type Element struct {
	value int
	pointerOnThisElement *Element
	pointerOnPtiviousElement *Element
}