package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	//Output:
	// Ace of Hearts
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}
