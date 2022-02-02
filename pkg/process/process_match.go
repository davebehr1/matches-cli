package process

import (
	"strconv"
	"strings"
)

func Process(matchResult string) (string, error) {
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
		return "", err
	}
	teamTwoScore, err := strconv.Atoi(strings.Trim(teamTwoResult[1], " "))
	if err != nil {
		return "", err
	}

	if teamOneScore > teamTwoScore {
		return teamOneResult[0], nil

	} else if teamTwoScore > teamOneScore {

		return teamTwoResult[0], nil

	} else if teamTwoScore == teamOneScore {
		return "draw", nil
	}
	return "", nil

}
