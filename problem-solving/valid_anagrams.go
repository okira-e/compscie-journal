func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters1 := make(map[byte]uint, len(s))
	letters2 := make(map[byte]uint, len(s))

	for _, val := range s {
		letters1[byte(val)] += 1
	}

	for _, val := range t {
		letters2[byte(val)] += 1
	}

	for key, _ := range letters1 {
		if letters1[key] != letters2[key] {
			return false
		}
	}

	for key, _ := range letters2 {
		if letters1[key] != letters2[key] {
			return false
		}
	}
	return true
}
