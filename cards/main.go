package main

func main() {
	cards := newDeck()

	hand, remainingCards := slicingFunc(cards, 5)
	hand.print()
	remainingCards.print()
}
