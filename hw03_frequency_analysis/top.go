package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Pair struct {
	Substring string
	Number    int
}
type PairList []Pair

func Top10(str string) []string {
	s := strings.Fields(str)

	valuesMap := make(map[string]int)
	for _, word := range s {
		valuesMap[word]++
	}

	pl := make(PairList, len(valuesMap))
	i := 0
	for k, v := range valuesMap {
		pl[i] = Pair{k, v}
		i++
	}

	sort.Slice(pl, func(i, j int) bool {
		if pl[i].Number == pl[j].Number {
			return pl[i].Substring < pl[j].Substring
		}
		return pl[i].Number > pl[j].Number
	})

	sortedSlice := make([]string, 0, len(pl))
	for _, v := range pl {
		sortedSlice = append(sortedSlice, v.Substring)
	}

	if len(sortedSlice) > 9 {
		return sortedSlice[0:10]
	}

	return sortedSlice
}
