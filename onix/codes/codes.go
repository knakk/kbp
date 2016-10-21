package codes

type Language int

const (
	English Language = iota
	Norwegian
)

type Item struct {
	Code, Label, Notes string
}
