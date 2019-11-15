//go:generate echo $PATH
package deck

import (
	"math/rand"
	"time"
)

// Rank represents the rank or value of a card
type Rank uint8

// Constants for card ranks
const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Suit represents the suit of a card
type Suit uint8

// Constants for card suits
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

var suits = [...]Suit{Club, Diamond, Heart, Spade}

// Card represents a single card
type Card struct {
	Rank
	Suit
}

// New creates ordered standard deck of 52 cards
func New() (deck []Card) {
	for _, suit := range suits {
		for rank := Ace; rank <= King; rank++ {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	return
}

//NewMulti returns a deck comprised of any number of standard decks
func NewMulti(n uint) (deck []Card) {
	for i := uint(0); i < n; i++ {
		deck = append(deck, New()...)
	}
	return
}

// Shuffle randomises the order of a deck
func Shuffle(d []Card) []Card {
	s := make([]Card, len(d))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(d))
	for i, p := range perm {
		s[i] = d[p]
	}
	return s
}
