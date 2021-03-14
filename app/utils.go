package app

import "net/url"

func GenerateHash(id int64) string {
	b62 := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	_hash := ""

	for id > 0 {
		m := id % 62
		id = id / 62
		_hash += b62[m]
	}

	return _hash
}

func ValidateURI(uri string) bool {
	_, err := url.Parse(uri)

	if err != nil {
		return false
	}

	return true
}
