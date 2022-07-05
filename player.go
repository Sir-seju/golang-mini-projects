package main

import (
	"fmt"
	"strconv"
	"time"
)

// Player - A struct defining attributes of a game player
type Player struct {
	Name      string // Name of the Player
	Purse     int    // Amount Player has to play the game
	Hand      []Card // Cards player has in hand during gameplay
	Ace       int    // whether player's hand has an ace
	HandValue int    // the value of a player's hand
	Blackjack bool   // whether player's hand is a blackjack
}

func (p *Player) setDefaults() {
	if p.Purse == 0 {
		p.Purse = 10_000
	}
	p.HandValue = 0
}

func (p *Player) deposit(sum int) {
	p.Purse += sum
}

func (p *Player) showBalance() {
	fmt.Printf("%s your balance is %d\n", p.Name, p.Purse)
}

func (p *Player) draw(deck *Deck) {
	x := []Card{deck.DrawCard()}
	for _, card := range x {
		p.HandValue += card.Value
		p.Hand = append(p.Hand, card)
		if card.Rank == "Ace" {
			p.Ace += 1
		}
		fmt.Printf("%s draws\n", p.Name)
	}
}

func (p *Player) showHand() {
	for _, card := range p.Hand {
		card.Show()
	}
	p.catchAce()
	if len(p.Hand) == 2 && p.Ace == 1 && p.HandValue == 21 {
		p.Blackjack = true
	}
}

func (p *Player) catchAce() {
	for i := 0; i < p.Ace; i++ {
		if p.HandValue > 21 {
			p.Ace -= 1
			p.HandValue -= 10
		}
	}
}

func (p *Player) declareHand() {
	fmt.Printf("Value of cards in %s's hand is %d\n", p.Name, p.HandValue)
}

func (p *Player) resetHand() {
	p.Hand = nil
	p.HandValue = 0
	p.Ace = 0
	p.Blackjack = false
}

func (p *Player) namePlayer() {
	var name string
	pprint("Hello Player! Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Enter a valid name in string alphabetical format")
	}
	p.Name = name
}

func (p *Player) bet() int {
	var input string
	var amount int
	p.showBalance()
	pprint("Place your bets: ")
	for {
		_, err := fmt.Scanln(&input)
		amount, err = strconv.Atoi(input)
		if err != nil {
			pprint("Enter a valid number!!")
			continue
		} else {
			validBet := amount <= p.Purse
			if !validBet {
				fmt.Printf("Your balance cannot fund this amount! choose an amount less than or equal to %d: ", p.Purse)
				continue
			}
			break
		}
	}
	p.Purse -= amount
	fmt.Printf("The amount %s bet is: %d\n", p.Name, amount)
	return amount
}

func (p *Player) hit(deck *Deck) {
	pprint("-----------------")
	p.draw(deck)
	p.showHand()
	p.declareHand()
	time.Sleep(2 * time.Second)
}

func (p *Player) stand() {
	pprint("-----------------")
	p.declareHand()
	time.Sleep(1 * time.Second)
}

func (p *Player) bust() bool {
	if p.HandValue > 21 {
		pprint("bust! You lose!")
		time.Sleep(2 * time.Second)
		return true
	}
	return false
}
