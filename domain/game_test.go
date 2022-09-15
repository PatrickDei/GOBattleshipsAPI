package domain

import "testing"

func TestDetermineGameStateReturnsPlayerTurnWhileGameIsInProgress(t *testing.T) {
	playerId := "1"
	opponentId := "2"
	g := Game{
		PlayerId:   playerId,
		OpponentId: opponentId,
		TurnCount:  0,
		Status:     InProgress,
	}

	gs := determineGameState(g)
	if gs.PlayerTurn == "" {
		t.Error("Game state didn't return players turn")
	}
	if gs.Won != "" {
		t.Error("Game state returned Won field while the match is still in progress")
	}
}

func TestDetermineOpponentIdReturnsAppropriateId(t *testing.T) {
	playerId := "1"
	opponentId := "2"
	g := Game{
		PlayerId:   playerId,
		OpponentId: opponentId,
	}

	oId := determineOpponentId(playerId, g)
	pId := determineOpponentId(opponentId, g)

	if oId != opponentId {
		t.Error("Method didn't return \"opponent_id\" which was expected on \"player_id\" input")
	}
	if pId != playerId {
		t.Error("Method didn't return \"player_id\" which was expected on \"opponent_id\" input")
	}
}
