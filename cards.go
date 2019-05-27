package main

import (
	"strconv"
)

func main() {
	// cards := deck{newCard(1), newCard(2), newCard(3)}
	// cards = append(cards, "Appended Card")

	newCards := newDeck()
	//hand, remainingCards := deal(newCards, 3)


	fileName := "cards.txt"
	newCards.saveToFile(fileName)

	//fmt.Println(ss)

	fromFile := newDeckFromFile(fileName)
	fromFile.print()


	//hand.print()
	//fmt.Println("-----------------------")
	//remainingCards.print()

	// cards.print()
	//otherCards.print()

}

func newCard(number int) string {
	return ("Card #" + strconv.Itoa(number))
}
