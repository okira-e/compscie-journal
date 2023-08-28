package main

import "strings"

func isValid(ascii uint8) bool {
	var isSmallLetter = func() bool {
		return ascii >= 65 && ascii <= 90
	}

	var isBigLetter = func() bool {
		return ascii >= 97 && ascii <= 122
	}

	var isDigit = func() bool {
		return ascii >= 48 && ascii <= 57
	}

	return isSmallLetter() || isBigLetter() || isDigit()
}

func isPalindrome(s string) bool {
	leftPointer := 0
	rightPointer := len(s) - 1
	for leftPointer <= rightPointer {
		if !isValid(s[leftPointer]) {
			leftPointer += 1
			continue
		}
		if !isValid(s[rightPointer]) {
			rightPointer -= 1
			continue
		}
		if strings.ToUpper(string(s[leftPointer])) != strings.ToUpper(string(s[rightPointer])) {
			return false
		}

		leftPointer += 1
		rightPointer -= 1
	}

	return true
}
