package deck

import (
	"testing"
)

// TestCreateCard tests basic card creation
func TestCreateCard(t *testing.T) {
	c := Card{Nine, Diamond}
	if c.Rank != Nine {
		t.Errorf("Card rank is incorrect, got: %d, want: %d.", c.Rank, Nine)
	}
	if c.Suit != Diamond {
		t.Errorf("Card suit is incorrect, got: %d, want: %d.", c.Suit, Diamond)
	}
}

//TestDeckLength tests basic deck creation
func TestDeckLength(t *testing.T) {
	deck := New()
	length := len(deck)
	if length != 52 {
		t.Errorf("Length of deck is incorrect, got: %d, want: %d", length, 52)
	}
}

//TestDeckDuplicates tests deck creation produces unique cards
func TestDeckDuplicates(t *testing.T) {
	deck := New()
	counts := make(map[Card]int)
	for _, card := range deck {
		counts[card]++
	}
	uniqueCards := len(counts)
	if uniqueCards != 52 {
		t.Errorf("Wrong number of unique cards, got: %d, want: %d", uniqueCards, 52)
	}
}

//TestShuffle tests deck shuffling
func TestShuffle(t *testing.T) {
	deck := New()
	Shuffle(deck)
}

//TestMulti tests creating more than one deck
func TestMulti(t *testing.T) {
	var n uint = 3
	deck := NewMulti(n)
	expLength := int(n * 52)
	length := len(deck)
	if length != expLength {
		t.Errorf("Length of deck is incorrect, got: %d, want: %d", length, expLength)
	}
}
