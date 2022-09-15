package domain

import (
	"strings"
	"testing"
)

var bf = BoardFactoryImpl{}

func TestGenerateNewBoardReturnsBoardWithFields(t *testing.T) {
	b := bf.GenerateNewBoard()

	if b.Fields == "" {
		t.Error("Generated board has no fields")
	}
}

func TestGenerateNewBoardReturnsBoardWithShips(t *testing.T) {
	b := bf.GenerateNewBoard()

	if b.ShipCount == 0 {
		t.Error("Generated board has ShipCount 0")
	}
	if !strings.Contains(b.Fields, Taken.String()) {
		t.Error("Generated board has no Ships")
	}
}
