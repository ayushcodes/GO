package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// create a new type of deck which is a slice of 'deck'
// which is slice of string

type deck []string

func (d deck) print() {
	for _ , card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for i := 1; i <= 9; i++ {
		for _, c := range cardSuits {
			cards = append(cards, strconv.Itoa(i) + " of " + c)
		}
		
	}
	return cards
	// typesOfCards := []string{"ace", "2", "3",}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize],  d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs,err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error in read file", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}