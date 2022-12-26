package main

type ElementOfDequeue struct {
	value int
	pointerOnThisElement *ElementOfDequeue
	pointerOnPtiviousElement *ElementOfDequeue
	pointerOnNextElement *ElementOfDequeue
}