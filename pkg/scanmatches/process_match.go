package scanmatches

import (
	"strconv"
	"strings"
)

func (rankTable *RankTable) Process(matchResult string) error {
	results := strings.Split(matchResult, ",")

	for key, result := range results {
		results[key] = strings.Trim(result, " ")
	}

	index := strings.LastIndex(results[0], " ")
	teamOneResult := make([]string, 2)

	teamOneResult[0] = results[0][:index]
	teamOneResult[1] = results[0][index:]

	index = strings.LastIndex(results[1], " ")
	teamTwoResult := make([]string, 2)

	teamTwoResult[0] = results[1][:index]
	teamTwoResult[1] = results[1][index:]

	teamOneScore, err := strconv.Atoi(strings.Trim(teamOneResult[1], " "))
	if err != nil {
		return err
	}
	teamTwoScore, err := strconv.Atoi(strings.Trim(teamTwoResult[1], " "))
	if err != nil {
		return err
	}

	if teamOneScore > teamTwoScore {

		rankTable.Table[teamOneResult[0]] += 3
		rankTable.Table[teamTwoResult[0]] += 0

	} else if teamTwoScore > teamOneScore {

		rankTable.Table[teamTwoResult[0]] += 3
		rankTable.Table[teamOneResult[0]] += 0

	} else if teamTwoScore == teamOneScore {
		rankTable.Table[teamOneResult[0]] += 1
		rankTable.Table[teamTwoResult[0]] += 1
	}
	return nil

}
