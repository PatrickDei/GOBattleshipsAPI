package domain

type Ship struct {
	Size         int
	Row          int
	Column       int
	IsHorizontal bool
}

func NewPatrolCraft() Ship {
	return Ship{Size: 1}
}

func NewSubmarine() Ship {
	return Ship{Size: 2}
}

func NewDestroyer() Ship {
	return Ship{Size: 3}
}

func NewBattleship() Ship {
	return Ship{Size: 4}
}
