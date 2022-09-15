package domain

type FieldStatus int

const (
	Unpopulated FieldStatus = iota
	Taken
	Hit
	Miss
)

func (f FieldStatus) String() string {
	return []string{".", "#", "X", "O"}[f]
}

func (f FieldStatus) Rune() rune {
	return []rune{'.', '#', 'X', 'O'}[f]
}
