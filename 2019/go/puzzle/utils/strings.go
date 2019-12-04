package utils

// StringRunes returns a slice of runes in a string
func StringRunes(str string) []rune {
	var res []rune

	for _, r := range str {
		res = append(res, r)
	}

	return res
}
