package writeranktable

import (
	"bufio"
	"fmt"
	"os"
)

//go:generate mockgen -destination=../mocks/writeranktable_mock.go -package=mocks . WriteRankTable
type WriteRankTable interface {
	WriteToFile(filePath string, rankTable string) error
}

type Writer struct{}

func NewWriter() Writer {
	return Writer{}
}

func (wr *Writer) WriteToFile(filePath string, rankTable string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(rankTable)
	if err != nil {
		return err
	}
	fmt.Printf("wrote rank table to %s\n", filePath)

	w.Flush()

	return nil
}
