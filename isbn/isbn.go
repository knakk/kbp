package isbn

import (
	"fmt"
	"strings"
)

// https://en.wikipedia.org/wiki/International_Standard_Book_Number
// https://en.wikipedia.org/wiki/Special:BookSources/979-8-6024-0545-3

type ISBN struct {
	Prefix     string
	Group      string
	Publisher  string
	Title      string
	CheckDigit string
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func validateISBN10(isbn []byte) bool {
	s, t := 0, 0
	for _, c := range isbn {
		t += int(c)
		s += t
	}
	return s%11 == 0
}

func validateISBN13(isbn [13]byte) bool {
	s := 0
	for i, c := range isbn {
		if i%2 == 0 {
			s += int(c)
		} else {
			s += 3 * int(c)
		}
	}
	return s%10 == 0
}

func numToChar(c rune) rune {
	if c == 10 {
		return 'X'
	}
	return c + '0'
}

func Parse(s string) (ISBN, error) {
	var res ISBN
	var n [13]byte
	i := 0
loop:
	for _, c := range s {
		switch c {
		case ' ', '-':
			continue
		case 'x', 'X':
			if i == 9 {
				n[9] = 10
				i++
				break loop
			} else {
				return ISBN{}, fmt.Errorf("isbn: unexpected character %q", string(c))
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n[i] = uint8(c - '0')
			i++
		default:
			return ISBN{}, fmt.Errorf("isbn: unexpected character %q", string(c))
		}
		if i == 13 {
			break
		}
	}

	switch i {
	case 10:
		if !validateISBN10(n[:10]) {
			return ISBN{}, fmt.Errorf("isbn: invalid checksum for %q", s)
		}
	case 13:
		if !validateISBN13(n) {
			return ISBN{}, fmt.Errorf("isbn: invalid checksum for %q", s)
		}
	default:
		// wrong number of digits
		return ISBN{}, fmt.Errorf("isbn: invalid: %q", s)
	}

	isbnstring := strings.Map(numToChar, string(n[:]))

	// Find prefix
	if i == 13 {
		// ISBN13
		res.Prefix = isbnstring[:3]
		isbnstring = isbnstring[3:]
		if res.Prefix != "978" && res.Prefix != "979" {
			return ISBN{}, fmt.Errorf("isbn: unknown prefix %q", res.Prefix)
		}
	} else {
		// ISBN10
		isbnstring = isbnstring[:10]
	}

	// Find registration group, length 1-5 digits
	for i := 5; i > 0; i-- {
		if group, ok := ranges[isbnstring[0:i]]; ok {
			res.Group = isbnstring[0:i]
			// Find registrant (publisher)
			for j := len(group.Ranges) - 1; j >= 0; j-- {
				l := len(group.Ranges[j][0])
				if i+l > len(isbnstring) {
					// should not be possible given isbnstring is length 10 or 13,
					// and none of the ranges ware more than 6 digits, but check just
					// in case to avoid any out of bounds panic
					continue
				}
				cand := isbnstring[i : i+l]
				if cand >= group.Ranges[j][0] && cand <= group.Ranges[j][1] {
					res.Publisher = cand
					res.Title = isbnstring[i+l : len(isbnstring)-1]
					res.CheckDigit = isbnstring[len(isbnstring)-1:]
					return res, nil
				}
			}
		}
	}

	// Sucessfully parsed ISBN are returned by now, so if we got this far
	// it means no matching isbn group and/or range.

	return res, fmt.Errorf("isbn: invalid: %q", s)
}

func (isbn ISBN) Hypenate() string {
	if isbn.Prefix != "" {
		return fmt.Sprintf("%s-%s-%s-%s-%s", isbn.Prefix, isbn.Group, isbn.Publisher, isbn.Title, isbn.CheckDigit)
	}
	return fmt.Sprintf("%s-%s-%s-%s", isbn.Group, isbn.Publisher, isbn.Title, isbn.CheckDigit)
}

// Prettify will attempt to parse the given string as an ISBN number, and output a
// hypenated representation. If the input is not a valid ISBN, it will be return unmodified.
func Prettify(s string) string {
	if isbn, err := Parse(s); err == nil {
		return isbn.Hypenate()
	}
	return s
}

var isbnCleaner = strings.NewReplacer("ISBN", "", "isbn", "", "-", "", " ", "")

func Clean(s string) string {
	return isbnCleaner.Replace(s)
}
