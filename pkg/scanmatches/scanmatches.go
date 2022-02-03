package scanmatches

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type StringReader interface {
	ReadString(delim byte) (string, error)
}

//go:generate mockgen -destination=../mocks/scanmatches_mock.go -package=mocks . ScanMatches
type ScanMatches interface {
	ScanFromFile(filePath string) ([]*TeamRank, error)
	ScanFromStdin(reader StringReader) ([]*TeamRank, error)
}

type RankTable struct {
	Table map[string]int
}

func NewRankTable() RankTable {
	return RankTable{Table: make(map[string]int)}
}

func (rankTable *RankTable) ScanFromFile(filePath string) ([]*TeamRank, error) {

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
		err := rankTable.Process(line)
		return nil, err

	}

	return rankTable.Sort(), nil
}

func (rankTable *RankTable) ScanFromStdin(reader StringReader) ([]*TeamRank, error) {
	for {
		fmt.Printf("Enter Match Result: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(text) == "done" {
			break
		} else {
			err := rankTable.Process(strings.TrimSpace(text))
			if err != nil {
				return nil, err
			}
		}
	}
	return rankTable.Sort(), nil
}
