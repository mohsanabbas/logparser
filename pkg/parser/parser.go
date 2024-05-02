package parser

import (
	"log-parser/pkg/logreader"
	"strings"
)

type GameParser interface {
	ParseGames(entries <-chan logreader.LogEntry) ([]GameInterface, error)
}

type quakeGameParser struct {
	parser KillParser
}

func NewParser(parser KillParser) GameParser {
	return &quakeGameParser{
		parser: parser,
	}
}

func (qgp quakeGameParser) ParseGames(entries <-chan logreader.LogEntry) ([]GameInterface, error) {
	var games []GameInterface
	var currentGame *Game

	for entry := range entries {
		if entry.Err != nil {
			return nil, entry.Err
		}

		line := entry.Line
		switch {
		case strings.Contains(line, "InitGame:"):
			if currentGame != nil {
				games = append(games, currentGame)
			}
			currentGame = &Game{
				ID:           len(games) + 1,
				Players:      make(map[string]*Player),
				KillsByMeans: make(map[string]int),
			}
		case strings.Contains(line, "Kill:"):
			if currentGame != nil {
				qgp.parser.ParseKill(line, currentGame)
			}
		}
	}

	if currentGame != nil {
		games = append(games, currentGame) // Append any remaining game not yet added
	}

	return games, nil
}
