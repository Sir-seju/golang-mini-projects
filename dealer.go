package main

import "time"

type Dealer struct {
	Player
}

func (d *Dealer) stand(deck *Deck) {
	for d.HandValue < 17 {
		d.draw(deck)
		d.showHand()
		time.Sleep(1 * time.Second)
	}
	d.declareHand()
}

func (d *Dealer) setDefaults() {
	d.Name = "dealer"
}
