package models

import "github.com/google/uuid"

type Match struct {
	ID        string     `json:"id"`
	GameBoard *GameBoard `json:"game_board"`
	Players   *[]Player  `json:"players"`
}

func NewMatch() *Match {
	return &Match{
		ID:        uuid.New().String(),
		GameBoard: NewGameBoard(),
		Players:   GetInitialPlayers(),
	}
}
