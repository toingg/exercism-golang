package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var number int
	switch card {
	case "ace":
		number = 11
	case "jack":
		number = 10
	case "queen":
		number = 10
	case "king":
		number = 10
	case "ten":
		number = 10
	case "two":
		number = 2
	case "three":
		number = 3
	case "four":
		number = 4
	case "five":
		number = 5
	case "six":
		number = 6
	case "seven":
		number = 7
	case "eight":
		number = 8
	case "nine":
		number = 9
	default:
		number = 0
	}
	return number
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumCards := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	var result string
	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case sumCards >= 12 && sumCards <= 16 && dealerCardValue < 7:
		result = "S"
	case sumCards == 21 && (dealerCardValue == 10 || dealerCardValue == 11):
		result = "S"
	case sumCards == 21 && (dealerCardValue != 10 || dealerCardValue == 11):
		result = "W"
	case sumCards >= 17 && sumCards <= 20:
		result = "S"
	case sumCards >= 12 && sumCards <= 16:
		if dealerCardValue >= 7 {
			result = "H"
		} else {
			result = "S"
		}
	case sumCards <= 11:
		result = "H"
	}

	return result
}
