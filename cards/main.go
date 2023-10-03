package main

func main() {
	//cards := newDeck()
	/* =======IT USED FOR SLICING======= */
	//
	//hand, remainingCards := slicingFunc(cards, 5)
	//hand.print()
	//remainingCards.print()

	/* =======IT USED FOR CHANGE TO STRING======= */
	//fmt.Println(cards.toString())

	/* =======IT USED FOR Write To File Function======= */
	//cards.saveToFile("saveToFile_cards")

	/* =======IT USED FOR Read the from the File======= */
	readCardsFromTheFile := newDeckFromFile("saveToFile_cards")
	readCardsFromTheFile.print()
}
