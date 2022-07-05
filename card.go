package main

import "fmt"

type Card struct {
	Suit  string
	Rank  string
	Value int
}

func (c Card) Show() {
	fmt.Printf("%s of %s\n", c.Rank, c.Suit)
}
