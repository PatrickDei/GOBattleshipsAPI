package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-battleships/dto"
	"go-battleships/service"
	"net/http"
)

const gameLocation = "game/match-"
const OpponentPathParam = "opponent_id"

type GameHandler struct {
	Facade service.GameFacade
}

func (gh GameHandler) ChallengePlayer(w http.ResponseWriter, r *http.Request) {
	opponentId := mux.Vars(r)[OpponentPathParam]

	var gc dto.GameCommand
	err := json.NewDecoder(r.Body).Decode(&gc)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		game, err := gh.Facade.ChallengeOpponent(gc.PlayerId, opponentId)
		if err != nil {
			writeResponse(w, err.Code, err.AsResponseMessage())
		} else {
			w.Header().Add("Location", gameLocation+game.Id)
			writeResponse(w, http.StatusCreated, game)
		}
	}
}