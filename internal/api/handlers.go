package api

import (
	"encoding/json"
	"net/http"
	"time"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	Unimplemented
}

func (s *Server) GameCreate(w http.ResponseWriter, r *http.Request, params GameCreateParams) {
	// 1. Прочитать тело
	var body NewGameRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		code := 400
		msg := "invalid JSON body"
		writeJSON(w, http.StatusBadRequest, ErrorBadRequest{
			StatusCode:   &code,
			ErrorMessage: &msg,
		})
		return
	}

	// 2. Простая валидация
	if body.Mode == "" {
		code := 400
		msg := "mode cannot be empty"
		writeJSON(w, http.StatusBadRequest, ErrorBadRequest{
			StatusCode:   &code,
			ErrorMessage: &msg,
		})
		return
	}

	// 3. Заглушка "создать игру"
	now := time.Now().UTC()
	game := Game{
		Id:          NewULID(),
		Mode:        GameMode(body.Mode),
		Rated:       body.Rated,
		Ended:       false,
		WhiteUserId: NewULID(),
		BlackUserId: NewULID(),
		StartAt:     &now,
	}

	// 4. Отправить ответ
	writeJSON(w, http.StatusCreated, game)
}

func (s *Server) Health(w http.ResponseWriter, r *http.Request, params HealthParams) {
	var body HealthParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		code := 400
		msg := "invalid JSON body"
		writeJSON(w, http.StatusBadRequest, ErrorBadRequest{
			StatusCode:   &code,
			ErrorMessage: &msg,
		})
		return
	}

	ok := OkResponse{
		Ok: true,
	}
	writeJSON(w, http.StatusCreated, ok)
}
