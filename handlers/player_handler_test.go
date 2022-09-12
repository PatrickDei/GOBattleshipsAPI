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

var mockService *service.MockPlayerService
var r *mux.Router
var ph PlayerHandler

var pc = dto.PlayerCommand{
	Name:  "some",
	Email: "mail",
}

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockPlayerService(ctrl)

	ph = PlayerHandler{Service: mockService}

	r = mux.NewRouter()
	r.HandleFunc("/player", ph.CreatePlayer).Methods(http.MethodPost)

	return func() {
		r = nil
		defer ctrl.Finish()
	}
}

func createPlayerRequest(pc dto.PlayerCommand) *http.Request {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(pc)
	request, _ := http.NewRequest(http.MethodPost, "/player", &buf)

	return request
}

func TestShouldReturnCreatedStatusWithLocationHeader(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	pdto := dto.PlayerDTO{
		Id:    "1",
		Name:  "some",
		Email: "mail",
	}

	mockService.EXPECT().CreatePlayer(pc).Return(&pdto, nil)

	request := createPlayerRequest(pc)

	// Act
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Didn't return 201 Created")
	}
	if recorder.Header().Get("Location") == "" {
		t.Error("Didn't return Location header")
	}
}

func TestShouldReturnConflictStatusWhenUsernameIsTaken(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().CreatePlayer(pc).Return(nil, errors.NewConflictError(errors.NewErrorBody("code", "arg")))

	request := createPlayerRequest(pc)

	// Act
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusConflict {
		t.Error("Didn't return 409 Conflict")
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
