package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"tic-tac-toe/internal/models"
)

func getDataPath() string {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "."
	}
	return filepath.Join(dataDir, "matches.json")
}

func SaveMatch(g *models.Match) {
	matches, err := LoadMatches()
	if err != nil {
		matches = &[]models.Match{}
	}

	oldMatch := false
	for i, match := range *matches {
		if match.ID == g.ID {
			(*matches)[i] = *g
			oldMatch = true
			break
		}
	}

	if !oldMatch {
		*matches = append(*matches, *g)
	}
	data, err := json.Marshal(matches)
	if err != nil {
		return
	}

	err = os.WriteFile(getDataPath(), data, 0644)
	if err != nil {
		return
	}
}

func LoadMatches() (*[]models.Match, error) {
	file, err := os.ReadFile(getDataPath())
	if err != nil {
		return nil, err
	}

	var matches []models.Match
	err = json.Unmarshal(file, &matches)
	if err != nil {
		return nil, err
	}
	return &matches, nil
}
