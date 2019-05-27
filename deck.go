package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Work with a bunch of cards

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"suit 1", "suit 2", "suit 3"}
	cardValues := []string{"Ace", "Cold", "Fire"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)

	}
	//string(bs) // Ace of suit, Two of suit 2
	s := strings.Split(string(bs), ",")
	return deck(s)
}
