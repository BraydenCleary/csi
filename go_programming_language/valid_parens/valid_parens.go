package main

import "unicode/utf8"

func main() {
}

var matchOpenClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isOpeningCharacter(r rune) (isOpening bool) {
	for key, _ := range matchOpenClose {
		if key == string(r) {
			isOpening = true
			break
		}
	}
	return
}

func isClosingCharacter(r rune) (isClosing bool) {
	for _, value := range matchOpenClose {
		if value == string(r) {
			isClosing = true
			break
		}
	}
	return
}

func isValid(input string) (output bool) {
	if utf8.RuneCountInString(input) == 1 {
		return
	}

	stack := []rune{}
	for _, r := range input {
		if isOpeningCharacter(r) {
			stack = append(stack, r)
		}

		if isClosingCharacter(r) {
			if len(stack) == 0 {
				return
			}

			popValue := stack[len(stack)-1]
			if matchOpenClose[string(popValue)] != string(r) {
				return
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		output = true
	}

	return
}
