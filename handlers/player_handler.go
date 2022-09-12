package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-battleships/dto"
	"go-battleships/service"
	"net/http"
)

const playerLocation = "/player/player-"
const PlayerPathParam = "player_id"

type PlayerHandler struct {
	Service service.PlayerService
}

func (ph PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var pc dto.PlayerCommand
	err := json.NewDecoder(r.Body).Decode(&pc)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		player, err := ph.Service.CreatePlayer(pc)
		if err != nil {
			writeResponse(w, err.Code, err.AsResponseMessage())
		} else {
			w.Header().Add("Location", playerLocation+player.Id)
			writeResponse(w, http.StatusCreated, nil)
		}
	}
}

func (ph PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)[PlayerPathParam]

	player, err := ph.Service.GetById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsResponseMessage())
	} else {
		writeResponse(w, http.StatusOK, player)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
