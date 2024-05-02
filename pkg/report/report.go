package report

import (
	"encoding/json"
	"fmt"
	"log-parser/pkg/parser"
	"strings"
)

func GenerateReport(games []parser.GameInterface) {
	for i, game := range games {
		uniquePlayers := make(map[string]bool)
		var playerList []string

		gameData := map[string]interface{}{
			"total_kills":    game.GetTotalKills(),
			"players":        []string{},
			"kills":          map[string]int{},
			"kills_by_means": game.GetKillsByMeans(),
		}

		for _, player := range game.GetPlayers() {
			playerName := cleanPlayerName(player.GetName())

			if playerName == "<world>" {
				continue // Skip adding <world> to the player list
			}
			if _, exists := uniquePlayers[playerName]; !exists {
				uniquePlayers[playerName] = true
				playerList = append(playerList, playerName)
				gameData["kills"].(map[string]int)[playerName] = player.GetKills()
			}
		}

		gameData["players"] = playerList
		gameJSON, _ := json.MarshalIndent(gameData, "", "  ")
		fmt.Printf("game-%d: %s\n", i+1, string(gameJSON))
	}
}

func cleanPlayerName(fullIdentifier string) string {
	parts := strings.Split(fullIdentifier, ": ")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}
	return fullIdentifier
}
