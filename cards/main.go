package main

import "fmt"

var a int
// a= 10
func main(){
	// var card string = "Ace Of Spades"
	card := newCard()
	// card := "sajgdcdsjhcbdsc"
	// card = 90
	a= 10
	fmt.Println(card)

	cards := deck{"a","b", "c"}
	cards = append(cards, "d")

	for i, cardd :=range cards {
		fmt.Println(i, cardd)
	}
	cards = newDeck()
	cards.print()
	hands, remainingDeck := deal(cards, 10)
	fmt.Println(hands)
	fmt.Println(remainingDeck)
	hands.print()
	a := "Hi there!"
	fmt.Println([]byte(a))
	fmt.Println(cards.toString())
	cards.saveToFile("mycards")
	data := newDeckFromFile("mycards")
	data.print()
	fmt.Println("##############")
	data.shuffle()
	data.print()

}

func newCard() string {
	return "Five of diamonds"
}