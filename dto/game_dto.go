package dto

type GameDTO struct {
	Id         string `json:"game_id"`
	PlayerId   string `json:"player_id"`
	OpponentId string `json:"opponent_id"`
	Starting   string `json:"starting"`
}
