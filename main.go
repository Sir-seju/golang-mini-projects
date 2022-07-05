package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// Start Game
	clear()
	pprint("Welcome to Uwasan's BlackJack Game!\n")
	time.Sleep(2 * time.Second)
	gameOn := true
	for gameOn {

		// Build and shuffle the card deck
		deck := Deck{}
		deck.AutoBuild()
		deck.Shuffle()

		// Instantiate player and dealer objects
		player := Player{}
		dealer := Dealer{}
		player.setDefaults()
		dealer.setDefaults()
		player.namePlayer()

		time.Sleep(1 * time.Second)
		fmt.Printf("Hello %s!!\n", player.Name)
		time.Sleep(1 * time.Second)
		fmt.Printf("The amount of money you have to bet in this game is %d\nGoodluck and bet wisely!\n", player.Purse)
		time.Sleep(3 * time.Second)

		// Start Game
		startGame := true
		for startGame {
			inGame := true
			for inGame {
				clear()
				player.resetHand()
				dealer.resetHand()
				bet := player.bet()
				time.Sleep(1 * time.Second)
				pprint("---------------------------------------------")
				dealer.draw(&deck)
				time.Sleep(1 * time.Second)
				dealer.showHand()
				time.Sleep(1 * time.Second)
				pprint("---------------------------------------------")
				player.draw(&deck)
				time.Sleep(1 * time.Second)
				player.draw(&deck)
				time.Sleep(1 * time.Second)
				player.showHand()
				pprint("---------------------------------------------")
				player.declareHand()
				if player.Blackjack {
					gameEval(&player, &dealer, bet)
				}
				time.Sleep(2 * time.Second)
				choice := hitOrStand()
				if choice == "H" {
					player.hit(&deck)
					if player.bust() {
						player.showBalance()
					} else {
						choice := hitOrStand()
						if choice == "H" {
							player.hit(&deck)
							if player.bust() {
								player.showBalance()
							} else {
								choice := hitOrStand()
								if choice == "H" {
									player.hit(&deck)
									if player.bust() {
										player.showBalance()
									} else {
										player.stand()
										dealer.stand(&deck)
										gameEval(&player, &dealer, bet)
									}
								} else {
									player.stand()
									dealer.stand(&deck)
									gameEval(&player, &dealer, bet)
								}
							}
						} else {
							player.stand()
							dealer.stand(&deck)
							gameEval(&player, &dealer, bet)
						}
					}
				} else {
					player.stand()
					dealer.stand(&deck)
					gameEval(&player, &dealer, bet)
				}
				if player.Purse < 100 {
					pprint("Game over!")
					inGame = false
					break
				}
				if playAgain() {
					continue
				} else {
					inGame = false
				}
			}
			startGame = false
		}
		gameOn = false
	}
	pprint("Thanks for playing my blackjack game")
	pprint("---------------------------------------------\n")
}

func pprint(p string) {
	fmt.Println(p)
}

func hitOrStand() string {
	pprint("Do you want to 'hit' or 'stand' h/s?: ")
	choice := ""
	for {
		_, err := fmt.Scanln(&choice)
		if err != nil {
			pprint("Invalid Input!!!")
			continue
		}
		choice = strings.ToUpper(choice)
		if choice == "H" || choice == "S" {
			return choice
		} else {
			fmt.Println(choice, "is an invalid option! choose 'h' or 's': ")
			continue
		}
	}
}

func bjCheck(p *Player, d *Dealer) bool {
	if p.Blackjack && d.HandValue < 21 {
		return true
	}
	return false
}

func winCheck(p *Player, d *Dealer) bool {
	if p.HandValue > d.HandValue && p.HandValue < 22 {
		return true
	} else if d.HandValue > 21 && p.HandValue < 22 {
		return true
	}
	return false
}

func drawCheck(p *Player, d *Dealer) bool {
	if p.Blackjack && d.Blackjack {
		return true
	}
	return p.HandValue == d.HandValue && d.HandValue < 22
}

func gameEval(p *Player, d *Dealer, bet int) {
	for {
		if d.HandValue > 21 {
			pprint("bust for dealer!")
			time.Sleep(1 * time.Second)
			fmt.Printf("Congratulations %s! you win!\n", p.Name)
			p.Purse += bet * 2
			p.showBalance()
			time.Sleep(3 * time.Second)
			break
		} else if bjCheck(p, d) {
			pprint("Blackjack!! you win\n")
			p.Purse += bet * 3
			p.showBalance()
			time.Sleep(3 * time.Second)
			break
		} else if winCheck(p, d) {
			fmt.Printf("Congratulations %s! you win!\n", p.Name)
			p.Purse += bet * 2
			p.showBalance()
			time.Sleep(3 * time.Second)
			break
		} else if drawCheck(p, d) {
			pprint("The game is a draw!\n")
			p.Purse += bet
			p.showBalance()
			time.Sleep(3 * time.Second)
			break
		} else {
			pprint("Sorry you lose!\n")
			time.Sleep(3 * time.Second)
			break
		}
	}
}

func playAgain() bool {
	var answer string
	pprint("Play again? \"y\" / \"n\": ")
	_, err := fmt.Scanln(&answer)
	if err != nil {
		pprint("Invalid Input!")
	}
	answer = strings.ToUpper(answer)
	if answer == "Y" {
		return true
	}
	return false
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
