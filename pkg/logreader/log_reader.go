package logreader

import (
	"bufio"
	"io"
)

const chunkSize = 100 // Buffer of 100 lines

type LogReader interface {
	Read(r io.Reader) ([]string, error)
	ReadLines(r io.Reader) <-chan LogEntry
}

type LogEntry struct {
	Line string
	Err  error
}

type fileLogReader struct{}

func (flr fileLogReader) ReadLines(r io.Reader) <-chan LogEntry {
	lines := make(chan LogEntry, chunkSize)
	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			lines <- LogEntry{
				Line: scanner.Text(),
				Err:  nil,
			}
		}
		if err := scanner.Err(); err != nil {
			lines <- LogEntry{
				Line: "",
				Err:  err,
			}
		}
	}()
	return lines
}

func NewFileLogReader() LogReader {
	return &fileLogReader{}
}

func (flr fileLogReader) Read(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
