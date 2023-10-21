package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()

	//cards := []string{"Ace of Diamond", newCard()}
	//cards := newDeckfromFile("cards.txt")
	cards.loadfromFile("cards.txt")
	fmt.Println("=====in main====")
	fmt.Println(cards)
	fmt.Println("=================")
	//cards = append(cards, "Six of Spades")
	//fmt.Println(cards)
	/*for i, card := range cards {ß
		fmt.Println(i, card)
	 }*/
	cards.print()
	cards.shuffle2()
	cards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("cards.txt")

	/*hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("salto")
	remainingCards.print()*/

}
