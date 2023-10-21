package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Espada", "Diamantes", "Corazones", "Treboles"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suits)
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
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func (d *deck) loadfromFile(filename string) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// mystring := string(bs)
	//*d = strings.Split(mystring, ",")
	*d = deck(strings.Split(string(bs), ","))
	//fmt.Println("=in loadfromFile=")
	//fmt.Println(*d)
	//fmt.Println("=================")
}

func newDeckfromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
		fmt.Println(i, newPosition)
	}
}

func (d deck) shuffle2() {
	//source := rand.NewSource(time.Now().UnixNano())
	// r := rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
		fmt.Println(i, newPosition)
	}

}
