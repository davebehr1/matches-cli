package scanmatches

import (
	"bufio"
	"log"
	"os"
)

type StringReader interface {
	ReadString(delim byte) (string, error)
}

//go:generate mockgen -destination=../mocks/scanmatches_mock.go -package=mocks . ScanMatches
type ScanMatches interface {
	ScanFromFile(filePath string) ([]TeamRank, error)
	ScanFromStdin(reader StringReader) (string, error)
}

func ReadMatch(r StringReader) (string, error) {
	match, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return match, nil
}

type RankTable struct {
	Table map[string]int
}

func NewRankTable() RankTable {
	return RankTable{Table: make(map[string]int)}
}

func (rankTable *RankTable) ScanFromFile(filePath string) ([]TeamRank, error) {

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, line := range text {
		rankTable.Process(line)

	}

	return rankTable.Sort(), nil
}

func (rankTable *RankTable) ScanFromStdin(reader StringReader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return text, nil
}