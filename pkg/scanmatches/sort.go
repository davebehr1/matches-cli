package scanmatches

import (
	"sort"
	"strings"
)

type TeamRank struct {
	Team string
	Rank int
}

func (rankTable *RankTable) Sort() []*TeamRank {

	var teamRanks []*TeamRank
	for k, v := range rankTable.Table {
		teamRanks = append(teamRanks, &TeamRank{k, v})
	}

	sort.Slice(teamRanks, func(i, j int) bool {
		return teamRanks[i].Rank > teamRanks[j].Rank
	})

	sort.Slice(teamRanks, func(i, j int) bool {
		if teamRanks[i].Rank == teamRanks[j].Rank {
			return strings.ToLower(teamRanks[i].Team[:1]) < strings.ToLower(teamRanks[j].Team[:1])
		}
		return false

	})

	return teamRanks
}
