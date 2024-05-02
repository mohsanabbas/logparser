package parser

import (
	"log-parser/pkg/logreader"
	"strings"
)

type GameParser interface {
	ParseGames(entries <-chan logreader.LogEntry, games chan<- GameInterface) error
}

type quakeGameParser struct {
	parser KillParser
}

func NewParser(parser KillParser) GameParser {
	return &quakeGameParser{
		parser: parser,
	}
}

func (qgp quakeGameParser) ParseGames(entries <-chan logreader.LogEntry, games chan<- GameInterface) error {
	var currentGame *Game

	for entry := range entries {
		if entry.Err != nil {
			return entry.Err
		}
		line := entry.Line
		switch {
		case strings.Contains(line, "InitGame:"):
			if currentGame != nil {
				games <- currentGame
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
		games <- currentGame
	}

	close(games)

	return nil
}
