package main

func main() {
	// var card string = "card"
	// cards := newDeck()

	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_card")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
