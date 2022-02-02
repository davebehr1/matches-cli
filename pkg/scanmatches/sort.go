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

	var ss []*TeamRank
	for k, v := range rankTable.Table {
		ss = append(ss, &TeamRank{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Rank > ss[j].Rank
	})

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Team == ss[j].Team {
			return strings.ToLower(ss[i].Team[:1]) < strings.ToLower(ss[j].Team[:1])
		}
		return false

	})

	return ss
}
