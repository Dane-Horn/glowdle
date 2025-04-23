package wordle

import "slices"

func validateWord(answer string, guess []key) ([]key, int) {
	answerRunes := []rune(answer)
	validatedKeys := make([]key, 5)
	validKeys := 0
	for i, k := range guess {
		if answerRunes[i] == k.value {
			k.state = correctKey
			validatedKeys[i] = k
			answerRunes[i] = -1
			validKeys++
		}
	}
	for i, k := range guess {
		if answerRunes[i] == -1 {
			continue
		}
		if slices.Contains(answerRunes, k.value) {
			k.state = presentKey
		} else {
			k.state = missingKey
		}
		validatedKeys[i] = k
	}
	return validatedKeys, validKeys
}
