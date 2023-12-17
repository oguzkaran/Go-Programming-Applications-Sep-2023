package card

import (
	"SampleGoLand/csd/util/algorithm"
	"fmt"
	"math/rand"
)

type Type int
type Value int

type Card struct {
	cardType  Type
	cardValue Value
}

const (
	Spade = iota
	Heart
	Diamond
	Club
)

const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Knave
	Queen
	King
	Ace
)

var cardTypes = [...]string{"Spade", "Heart", "Diamond", "Club"}
var cardValues = [...]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Knave", "Queen", "King", "Ace"}

const defaultCount = 100

func NewCard(cardType Type, cardValue Value) *Card {
	return &Card{cardType, cardValue}
}

func NewCardString(str string) *Card {
	card := Card{}

	card.SetCard(str)

	return &card
}

func NewDeck() []Card {
	deck := make([]Card, 52)
	i := 0

	for t := Spade; t <= Club; t++ {
		for v := Two; v <= Ace; v++ {
			deck[i] = Card{Type(t), Value(v)}
			i++
		}
	}

	return deck
}

func NewShuffledDeck() []Card {
	return NewShuffledCountDeck(defaultCount)
}

func NewShuffledCountDeck(count int) []Card {
	return ShuffleCount(NewDeck(), count)
}

func (c *Card) SetType(cardType Type) {
	if cardType < Spade || cardType > Club {
		panic("invalid card type!...")
	}

	c.cardType = cardType
}

func (c *Card) GetType() Type {
	return c.cardType
}

func (c *Card) SetValue(cardValue Value) {
	if cardValue < Two || cardValue > Ace {
		panic("invalid card value!...")
	}

	c.cardValue = cardValue
}

func (c *Card) GetValue() Value {
	return c.cardValue
}

func (c *Card) SetCard(str string) {
	//TODO: Kart bilgisi örneğin Club-King biçiminde bir yazı olarak verilecektir ve ilgili Card nesnesi bu yazıdan oluşturulacaktır
	// Yazının formatı bu şekilde olacaktır. Hatalı bir yazı için panic fonksiyonu çağrılacaktır
}

func ShuffleCount(deck []Card, count int) []Card {
	length := len(deck)

	for i := 0; i < count; i++ {
		algorithm.Swap(&deck[rand.Intn(length)], &deck[rand.Intn(length)])
	}

	return deck
}

func Shuffle(deck []Card) []Card {
	return ShuffleCount(deck, defaultCount)
}

func (c *Card) String() string {
	return fmt.Sprintf("%s-%s", cardTypes[c.cardType], cardValues[c.cardValue])
}
