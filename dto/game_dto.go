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
	opponentId := determineOpponentId(playerId, g)
	gs := determineGameState(g)
	return GameStateDTO{
		Player: boardState{
			PlayerId: playerId,
			Board:    b.GetFieldsAsSlice(),
		},
		Opponent: boardState{
			PlayerId: opponentId,
			Board:    domain.NewEmptyBoard().GetFieldsAsSlice(),
		},
		Game: gs,
	}
}

func determineOpponentId(playerId string, g domain.Game) string {
	var opponentId string
	if g.PlayerId == playerId {
		opponentId = g.OpponentId
	} else {
		opponentId = g.PlayerId
	}
	return opponentId
}

func determineGameState(g domain.Game) gameState {
	var gs gameState
	if g.Status == domain.InProgress {
		gs = gameState{
			PlayerTurn: g.DetermineIdOfPlayersTurn(),
		}
	} else if g.Status == domain.Finished {
		gs = gameState{
			Won: g.DetermineIdOfPlayersTurn(),
		}
	}

	return gs
}

type boardState struct {
	PlayerId string   `json:"player_id"`
	Board    []string `json:"board"`
}

type gameState struct {
	PlayerTurn string `json:"player_turn,omitempty"`
	Won        string `json:"won,omitempty"`
}
