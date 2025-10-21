package models

type GameBoard struct {
	Board [3][3]string `json:"board"`
}

func NewGameBoard() *GameBoard {
	return &GameBoard{
		Board: [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}
}
