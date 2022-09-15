package domain

type GameOutcome int

const (
	Won GameOutcome = iota
	Lost
	NotFinished
)

func (g GameOutcome) String() string {
	return []string{"WON", "LOST", "IN_PROGRESS"}[g]
}
