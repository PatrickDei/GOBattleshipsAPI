package domain

import (
	"math/rand"
	"strings"
	"time"
)

//go:generate mockgen -destination=../mocks/domain/mock_board_factory.go -package=domain -source=board_factory.go BoardFactory
type BoardFactory interface {
	PopulateBoard(*Board)

	PlacePatrolCraft(*Board) Ship
	PlaceSubmarine(*Board) Ship
	PlaceDestroyer(*Board) Ship
	PlaceBattleship(*Board) Ship
}

const patrolCraftAmount = 4
const submarineAmount = 3
const destroyerAmount = 2
const battleshipAmount = 1

type BoardFactoryImpl struct{}

func (bf BoardFactoryImpl) PopulateBoard(b *Board) {
	ships := make([]Ship, 0)
	for i := 0; i < battleshipAmount; i++ {
		ships = append(ships, bf.PlaceBattleship(b))
	}
	for i := 0; i < destroyerAmount; i++ {
		ships = append(ships, bf.PlaceDestroyer(b))
	}
	for i := 0; i < submarineAmount; i++ {
		ships = append(ships, bf.PlaceSubmarine(b))
	}
	for i := 0; i < patrolCraftAmount; i++ {
		ships = append(ships, bf.PlacePatrolCraft(b))
	}

	b.ShipCount = len(ships)
}

func (bf BoardFactoryImpl) PlacePatrolCraft(b *Board) Ship {
	pc := NewPatrolCraft()
	return bf.placeShip(b, &pc)
}

func (bf BoardFactoryImpl) PlaceSubmarine(b *Board) Ship {
	s := NewSubmarine()
	return bf.placeShip(b, &s)
}

func (bf BoardFactoryImpl) PlaceDestroyer(b *Board) Ship {
	d := NewDestroyer()
	return bf.placeShip(b, &d)
}

func (bf BoardFactoryImpl) PlaceBattleship(b *Board) Ship {
	bs := NewBattleship()
	return bf.placeShip(b, &bs)
}

func (bf BoardFactoryImpl) placeShip(b *Board, s *Ship) Ship {
	bf.positionShipOnAvailableSpot(b, s)

	bf.placeShipOnBoard(b, s)

	return *s
}

func (bf BoardFactoryImpl) positionShipOnAvailableSpot(b *Board, s *Ship) {
	var row, col int
	var horizontal bool

	for available := true; available; available = !bf.isSpotAvailable(b, s) {
		row, col, horizontal = bf.generateRandomSpot(s.Size)
		s.Row = row
		s.Column = col
		s.IsHorizontal = horizontal
	}
}

func (bf BoardFactoryImpl) generateRandomSpot(shipSize int) (int, int, bool) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	horizontal := r.Intn(2) == 1
	var row, col int
	if horizontal {
		row = r.Intn(BoardRowCount)
		col = r.Intn(BoardColumnCount - shipSize + 1)
	} else {
		row = r.Intn(BoardRowCount - shipSize + 1)
		col = r.Intn(BoardColumnCount)
	}

	return row, col, horizontal
}

func (bf BoardFactoryImpl) isSpotAvailable(b *Board, s *Ship) bool {
	isAvailable := true

	if s.IsHorizontal {
		for i := s.Column; i < s.Column+s.Size; i++ {
			if bf.areSpotsAroundPointTaken(b, s.Row, i) {
				isAvailable = false
				break
			}
		}
	}

	if !s.IsHorizontal {
		for i := s.Row; i < s.Row+s.Size; i++ {
			if bf.areSpotsAroundPointTaken(b, i, s.Column) {
				isAvailable = false
				break
			}
		}
	}

	return isAvailable
}

func (bf BoardFactoryImpl) areSpotsAroundPointTaken(b *Board, row int, col int) bool {
	return bf.isAboveTaken(b, row, col) ||
		bf.isBelowTaken(b, row, col) ||
		bf.isLeftTaken(b, row, col) ||
		bf.isRightTaken(b, row, col) ||
		string(b.GetFieldsAsSlice()[row][col]) != "."
}

func (bf BoardFactoryImpl) isAboveTaken(b *Board, row int, col int) bool {
	if row == 0 {
		return false
	}
	return string(b.GetFieldsAsSlice()[row-1][col]) != "." // todo make enum
}

func (bf BoardFactoryImpl) isBelowTaken(b *Board, row int, col int) bool {
	if row == BoardRowCount-1 {
		return false
	}
	return string(b.GetFieldsAsSlice()[row+1][col]) != "." // todo make enum
}

func (bf BoardFactoryImpl) isLeftTaken(b *Board, row int, col int) bool {
	if col == 0 {
		return false
	}
	return string(b.GetFieldsAsSlice()[row][col-1]) != "." // todo make enum
}

func (bf BoardFactoryImpl) isRightTaken(b *Board, row int, col int) bool {
	if col == BoardColumnCount-1 {
		return false
	}
	return string(b.GetFieldsAsSlice()[row][col+1]) != "." // todo make enum
}

func (bf BoardFactoryImpl) placeShipOnBoard(b *Board, s *Ship) {
	rows := b.GetFieldsAsSlice()

	for i := 0; i < s.Size; i++ {
		if s.IsHorizontal {
			row := []rune(strings.Clone(rows[s.Row]))
			row[s.Column+i] = '#'
			rows[s.Row] = string(row)
		} else {
			row := []rune(strings.Clone(rows[s.Row+i]))
			row[s.Column] = '#'
			rows[s.Row+i] = string(row)
		}
	}

	b.Fields = strings.Join(rows, "")
}

func NewBoardFactory() BoardFactory {
	return BoardFactoryImpl{}
}
