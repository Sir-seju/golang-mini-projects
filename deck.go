package main

import (
	"math/rand"
	"time"
)

// Deck - The deck of cards for the game
type Deck struct {
	Cards []Card
}

// Build - Build a deck of cards for the game
func (d *Deck) Build() {
	cardValues := map[string]int{
		"Ace": 11, "1": 1, "2": 2, "3": 3, "4": 4,
		"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"10": 10, "Jack": 10, "Queen": 10, "King": 10,
	}
	suitList := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	rankList := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range suitList {
		for _, rank := range rankList {
			value := cardValues[rank]
			d.Cards = append(d.Cards, Card{Rank: rank, Suit: suit, Value: value})
		}
	}
}

// Show - Print out the cards in the game deck
func (d *Deck) Show() {
	for _, card := range d.Cards {
		card.Show()
	}
}

// AutoBuild - Build the deck of cards twice for a full stack
func (d *Deck) AutoBuild() {
	d.Build()
	d.Build()
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// DrawCard - draw method from deck of cards
func (d *Deck) DrawCard() Card {
	i := len(d.Cards) - 1
	card := d.Cards[i]
	d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
	return card
}
