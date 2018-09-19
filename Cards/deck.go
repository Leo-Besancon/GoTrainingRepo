package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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
	return d[handSize:], d[:handSize]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveDeckToFile(filename string) {

	byteDeck := []byte(d.toString())

	ioutil.WriteFile(filename, byteDeck, 0666)
}

func newDeckFromFile(filename string) deck {

	byteDeck, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("FATAL ERROR : ", err)
		os.Exit(1)
	}

	s := string(byteDeck)

	d := deck(strings.Split(s, ","))

	return d
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]

	}

}
