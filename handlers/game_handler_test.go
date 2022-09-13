package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"go-battleships/dto"
	"go-battleships/errors"
	"go-battleships/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockFacade *service.MockGameFacade
var gameRouter *mux.Router
var gh GameHandler

var playerId = "1"
var opponentId = "2"
var gc = dto.GameCommand{PlayerId: playerId}

func gameSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockFacade = service.NewMockGameFacade(ctrl)

	gh = GameHandler{Facade: mockFacade}

	r = mux.NewRouter()
	r.HandleFunc("/player/{"+OpponentPathParam+"}/game", gh.ChallengePlayer).Methods(http.MethodPost)

	return func() {
		r = nil
		defer ctrl.Finish()
	}
}

func createGameRequest(gc dto.GameCommand) *http.Request {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(gc)
	request, _ := http.NewRequest(http.MethodPost, "/player/{"+opponentId+"}/game", &buf)

	return request
}

func TestChallengePlayerShouldReturnCreatedGameWithLocationHeader(t *testing.T) {
	teardown := gameSetup(t)
	defer teardown()

	gdto := dto.GameDTO{
		Id:         "1",
		PlayerId:   playerId,
		OpponentId: opponentId,
		Starting:   "1",
	}

	mockFacade.EXPECT().ChallengeOpponent(gc.PlayerId, gomock.Any()).Return(&gdto, nil)

	request := createGameRequest(gc)

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Error("Returned status is not created")
	}
	if recorder.Header().Get("Location") == "" {
		t.Error("Didn't return Location header")
	}
}

func TestChallengePlayerShouldReturnNotFoundWhenPlayerDoesntExist(t *testing.T) {
	teardown := gameSetup(t)
	defer teardown()

	mockFacade.EXPECT().ChallengeOpponent(gc.PlayerId, gomock.Any()).Return(nil, errors.NewNotFoundError(errors.NewErrorBody("code", "arg")))

	request := createGameRequest(gc)

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Error("Didn't return 404 Not Found")
	}

	var body errors.AppError
	_ = json.NewDecoder(recorder.Body).Decode(&body)
	if _, ok := body.Body[errors.ErrorBodyErrorCode]; ok {
		t.Error("Didn't return an error code")
	}
	if _, ok := body.Body[errors.ErrorBodyErrorArg]; ok {
		t.Error("Didn't return an error argument")
	}
}
