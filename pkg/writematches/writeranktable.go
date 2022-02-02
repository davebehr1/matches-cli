package writematches

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
	n4, err := w.WriteString(rankTable)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

	return nil
}
