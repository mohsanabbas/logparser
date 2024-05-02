package main

import (
	"log"
	"log-parser/pkg/logreader"
	"log-parser/pkg/parser"
	"log-parser/pkg/report"
	"os"
	"path/filepath"
)

func main() {
	absPath, err := filepath.Abs("data/games.log")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logR := logreader.NewFileLogReader()
	logEntries := logR.ReadLines(file)

	kp := parser.NewKillParser()
	qgp := parser.NewParser(kp)

	gamesChan := make(chan parser.GameInterface)
	go func() {
		if err := qgp.ParseGames(logEntries, gamesChan); err != nil {
			log.Fatal(err)
		}
	}()

	// Generate the report
	report.GenerateReport(gamesChan)
}
