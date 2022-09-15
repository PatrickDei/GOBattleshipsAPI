package dto

type GameDTO struct {
	Id         string `json:"game_id"`
	PlayerId   string `json:"player_id"`
	OpponentId string `json:"opponent_id"`
	Starting   string `json:"starting"`
}

type GameStateDTO struct {
	Player   BoardState `json:"self"`
	Opponent BoardState `json:"opponent"`
	Game     GameState  `json:"game"`
}

type BoardState struct {
	PlayerId string   `json:"player_id"`
	Board    []string `json:"board"`
}

type GameState struct {
	PlayerTurn string `json:"player_turn,omitempty"`
	Won        string `json:"won,omitempty"`
}
