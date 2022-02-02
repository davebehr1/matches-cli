package process

import (
	"strconv"
	"strings"
)

func Process(matchResult string) (string, error) {
	teamOneResult := make([]string, 2)
	teamTwoResult := make([]string, 2)
	results := strings.Split(matchResult, ",")

	teamOneResult = strings.Split(results[0], " ")
	teamTwoResult = strings.Split(results[0], " ")

	teamOneScore, err := strconv.Atoi(strings.Trim(teamOneResult[1], " "))
	if err != nil {
		return "", err
	}
	teamTwoScore, err := strconv.Atoi(strings.Trim(teamTwoResult[1], " "))

	if teamOneScore > teamTwoScore {
		return teamOneResult[0], nil

	} else if teamTwoScore > teamOneScore {

		return teamTwoResult[0], nil

	} else if teamTwoScore == teamOneScore {
		return "draw", nil
	}
	return "", nil

}
