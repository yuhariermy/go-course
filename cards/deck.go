package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

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

// For Slicing
func slicingFunc(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// For change slicing to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to file after run it, (d deck) was receiver. it could be use cards.saveToFile
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is default err, for can write and read
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}
