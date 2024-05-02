package parser

import (
	"log"
	"strings"
)

type KillParser interface {
	ParseKill(line string, game *Game)
}

type quakeKillParser struct{}

func NewKillParser() KillParser {
	return &quakeKillParser{}
}
func (qkp quakeKillParser) ParseKill(line string, game *Game) {
	parts := strings.Split(line, "Kill:")
	if len(parts) < 2 {
		log.Printf("Invalid line: %s", line)
		return
	}
	details := strings.Split(parts[1], " by ")
	if len(details) < 2 {
		log.Printf("Invalid line: %s", line)
		return
	}

	killerAndVictim := strings.Split(details[0], " killed ")
	if len(killerAndVictim) < 2 {
		log.Printf("Invalid line: %s", line)
		return
	}
	killer := strings.TrimSpace(killerAndVictim[0])
	victim := strings.TrimSpace(killerAndVictim[1])

	means := strings.TrimSpace(details[1])

	if killer == "<world>" {
		// Decrement victim's kills
		if victimPlayer, exists := game.Players[victim]; exists {
			victimPlayer.Kills--
			if victimPlayer.Kills < 0 {
				victimPlayer.Kills = 0
			}
		}
	}
	// Increment killer's kills
	if _, exists := game.Players[killer]; !exists {
		game.Players[killer] = &Player{Name: killer, Kills: 0}
	}
	game.Players[killer].Kills++
	game.TotalKills++

	// Record death means
	if _, exists := game.KillsByMeans[means]; !exists {
		game.KillsByMeans[means] = 0
	}
	game.KillsByMeans[means]++
}
