package models

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func GetInitialPlayers() ([]Player, error) {
	Players := []Player{
		{1, "Player 1", "X"},
		{2, "Player 2", "O"},
	}

	return Players, nil
}
