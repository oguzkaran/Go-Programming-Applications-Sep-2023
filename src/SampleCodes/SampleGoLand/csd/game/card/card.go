package card

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
	Clubs
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

func New(cardType Type, cardValue Value) *Card {
	return &Card{cardType, cardValue}
}

func NewDeck() []Card {
	deck := make([]Card, 52)
	i := 0

	for t := Spade; t <= Clubs; t++ {
		for v := Two; v <= Ace; v++ {
			deck[i] = Card{Type(t), Value(v)}
			i++
		}
	}

	return deck
}

func NewShuffledDeck() []Card {
	
}

func (c *Card) SetType(cardType Type) {
	if cardType < Spade || cardType > Clubs {
		panic("invalid card type!...")
	}

	c.cardType = cardType
}

func (c *Card) GetType() Type {
	return c.cardType
}

func (c *Card) SetValue(cardValue Value) {
	if cardValue < Two || cardValue > Ace {
		panic("invalid card type!...")
	}

	c.cardValue = cardValue
}

func (c *Card) GetValue() Value {
	return c.cardValue
}
