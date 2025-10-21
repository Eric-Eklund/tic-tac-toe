package models

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func GetInitialPlayers() ([]Player, error) {
	Players := []Player{
		{1, "Eric", "X"},
		{2, "Jenny", "O"},
	}

	return Players, nil
}
