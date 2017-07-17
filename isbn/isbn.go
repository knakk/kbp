package isbn

import "fmt"

type ISBN [13]byte

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func Parse(s string) (ISBN, error) {
	var n ISBN
	i := 0
	for _, c := range s {
		if c == '-' {
			continue
		}
		if isDigit(c) {
			n[i] = uint8(c - '0')
			i++
			if i == 11 {
				break
			}
		}
		if i == 10 && c == 'x' || c == 'X' {
			n[9] = 10
			break
		}
	}
	if i != 10 && i != 13 {
		return n, fmt.Errorf("incorrect number of digits: %d != 10|13", i)
	}
	return ISBN{}, nil
}

func (n ISBN) Is10() bool {
	return n[11] == 0 &&
		n[10] == 0 &&
		n[9] == 0
}

func (n ISBN) Is13() bool { return !n.Is10() }

func (n ISBN) Hyphenate() string {
	return "TODO"
}
