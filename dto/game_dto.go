package dto

import "go-battleships/domain"

type GameDTO struct {
	Id         string `json:"game_id"`
	PlayerId   string `json:"player_id"`
	OpponentId string `json:"opponent_id"`
	Starting   string `json:"starting"`
}

type GameStateDTO struct {
	Player   boardState `json:"self"`
	Opponent boardState `json:"opponent"`
	Game     gameState  `json:"game"`
}

func NewGameStateDTO(playerId string, g domain.Game, b domain.Board) GameStateDTO {
	var opponentId string
	if g.PlayerId == playerId {
		opponentId = g.OpponentId
	} else {
		opponentId = g.PlayerId
	}
	return GameStateDTO{
		Player: boardState{
			PlayerId: playerId,
			Board:    b.GetFieldsAsSlice(),
		},
		Opponent: boardState{
			PlayerId: opponentId,
			Board:    domain.NewEmptyBoard().GetFieldsAsSlice(),
		},
		Game: gameState{
			PlayerTurn: g.DetermineIdOfPlayersTurn(),
		},
	}
}

type boardState struct {
	PlayerId string   `json:"player_id"`
	Board    []string `json:"board"`
}

type gameState struct {
	PlayerTurn string `json:"player_turn"`
}
