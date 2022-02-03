package acceptance

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/davebehr1/spanassessment/pkg/writeranktable"
	. "github.com/onsi/gomega"
)

func TestGenerateRankTableCli(t *testing.T) {
	t.Run("Run cli with file input and stdout output", func(t *testing.T) {
		g := NewGomegaWithT(t)
		scanner := scanmatches.NewRankTable()
		writer := writeranktable.NewWriter()

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(&scanner, &writer)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--f=../../matches.txt"})
		err := RootCommand.Execute()
		g.Expect(err).To(BeNil())

		result := bf.String()
		g.Expect(result).To(Equal("1. Lions, 11 pts \n2. Tarantulas, 6 pts \n3. FC Awesome, 1 pts \n4. Snakes, 1 pts \n5. Cheetahs, 0 pts \n6. Grouches, 0 pts \n7. Pumas, 0 pts \n1. Lions, 22 pts \n2. Tarantulas, 12 pts \n3. FC Awesome, 2 pts \n4. Snakes, 2 pts \n5. Cheetahs, 0 pts \n6. Grouches, 0 pts \n7. Pumas, 0 pts \n"))
	})
	t.Run("Run cli with file input and file output", func(t *testing.T) {
		g := NewGomegaWithT(t)
		outputFilepath := "ranktable.txt"
		scanner := scanmatches.NewRankTable()
		writer := writeranktable.NewWriter()

		RootCommand := cmd.Initialize(&scanner, &writer)
		RootCommand.SetArgs([]string{"grt", "--f=../../matches.txt", fmt.Sprintf("--o=%s", outputFilepath)})
		err := RootCommand.Execute()
		g.Expect(err).To(BeNil())

		file, err := os.Open(outputFilepath)
		g.Expect(err).To(BeNil())

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)
		var text []string

		for fileScanner.Scan() {
			text = append(text, fileScanner.Text())
		}

		file.Close()

		g.Expect(len(text)).To(Equal(7))

		t.Cleanup(func() {
			err := os.Remove(outputFilepath)
			g.Expect(err).To(BeNil())
		})
	})
	t.Run("Run cli with user input and stdout output", func(t *testing.T) {
		g := NewGomegaWithT(t)
		scanner := scanmatches.NewRankTable()
		writer := writeranktable.NewWriter()

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(&scanner, &writer)
		RootCommand.SetOut(bf)
		RootCommand.SetIn(bytes.NewBufferString("Lions 3, Cheetahs 0 \ndone\n"))
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt"})
		err := RootCommand.Execute()
		g.Expect(err).To(BeNil())

		result := bf.String()
		g.Expect(result).To(Equal("1. Lions, 3 pts \n2. Cheetahs, 0 pts \n"))
	})
	t.Run("Run cli with user input and file output", func(t *testing.T) {
		g := NewGomegaWithT(t)
		outputFilepath := "ranktable.txt"
		scanner := scanmatches.NewRankTable()
		writer := writeranktable.NewWriter()

		RootCommand := cmd.Initialize(&scanner, &writer)
		RootCommand.SetIn(bytes.NewBufferString("Lions 3, Cheetahs 0 \ndone\n"))
		RootCommand.SetArgs([]string{"grt", fmt.Sprintf("--o=%s", outputFilepath)})
		err := RootCommand.Execute()
		g.Expect(err).To(BeNil())

		file, err := os.Open(outputFilepath)
		g.Expect(err).To(BeNil())

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)
		var text []string

		for fileScanner.Scan() {
			text = append(text, fileScanner.Text())
		}

		file.Close()

		g.Expect(len(text)).To(Equal(2))

		t.Cleanup(func() {
			err := os.Remove(outputFilepath)
			g.Expect(err).To(BeNil())
		})
	})
}
