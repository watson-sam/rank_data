package rank_data

import (
	"sort"
)

type IndexedData struct {
	Value float64
	Index int

	Count float64
	Sum   float64
}

func setUpRankFunc(data []float64) ([]IndexedData, []float64) {
	indexedData := make([]IndexedData, len(data))
	ranks := make([]float64, len(data))
	for i, v := range data {
		indexedData[i] = IndexedData{Value: v, Index: i}
	}
	sort.Slice(indexedData, func(i, j int) bool {
		return indexedData[i].Value < indexedData[j].Value
	})
	return indexedData, ranks
}

func updatePrevious(
	count float64,
	sum float64,
	rank float64,
	i int,
	indexedData []IndexedData,
) (float64, float64, []IndexedData) {
	count += 1
	sum += rank
	for c := 0; c < int(count); c++ {
		indexedData[i-c].Sum = sum
		indexedData[i-c].Count = count
	}
	return 0, 0, indexedData
}

func RankAverage(data []float64) []float64 {
	indexedData, ranks := setUpRankFunc(data)
	var sum, count, rank float64 = 0, 0, 0

	for i, indexed := range indexedData {
		rank++
		if i < len(indexedData)-1 {
			if indexed.Value != indexedData[i+1].Value {
				if count > 0 {
					count, sum, indexedData = updatePrevious(count, sum, rank, i, indexedData)
				} else {
					indexedData[i].Sum = rank
					indexedData[i].Count = 1
				}
			} else {
				sum += rank
				count += 1
			}
		} else {
			if count > 0 {
				_, _, indexedData = updatePrevious(count, sum, rank, i, indexedData)
			} else {
				indexedData[i].Sum = (sum + rank)
				indexedData[i].Count = (count + 1)
			}
		}
	}
	for _, indexed := range indexedData {
		ranks[indexed.Index] = float64(indexed.Sum / indexed.Count)
	}

	return ranks
}

func RankMin(data []float64) []float64 {
	indexedData, ranks := setUpRankFunc(data)
	var rank float64 = 1
	var sameN float64 = 1
	for i, indexed := range indexedData {
		ranks[indexed.Index] = float64(rank)
		if i < len(indexedData)-1 && indexed.Value != indexedData[i+1].Value {
			rank += sameN
			sameN = 1
		} else {
			sameN++
		}
	}
	return ranks
}

func RankMax(data []float64) []float64 {
	indexedData, ranks := setUpRankFunc(data)
	sort.Slice(indexedData, func(i, j int) bool {
		return indexedData[i].Value > indexedData[j].Value
	})
	var rank float64 = float64(len(indexedData))
	var sameN float64 = 1
	for i, indexed := range indexedData {
		ranks[indexed.Index] = float64(rank)
		if i < len(indexedData)-1 && indexed.Value != indexedData[i+1].Value {
			rank -= sameN
			sameN = 1
		} else {
			sameN++
		}
	}
	return ranks
}
func RankDense(data []float64) []float64 {
	indexedData, ranks := setUpRankFunc(data)
	var rank float64 = 1
	for i, indexed := range indexedData {
		ranks[indexed.Index] = float64(rank)
		if i < len(indexedData)-1 && indexed.Value != indexedData[i+1].Value {
			rank++
		}
	}
	return ranks
}

func RankOrdinal(data []float64) []float64 {
	indexedData, ranks := setUpRankFunc(data)
	for i, indexed := range indexedData {
		ranks[indexed.Index] = float64(i + 1)
	}
	return ranks
}

// RankData Assign ranks to data, dealing with ties appropriately.
// Ranks begin at 1.  The `method` argument controls how ranks are assigned
// to equal values.
//
// Parameters:
// a: []float64 The array of values to be ranked.
// method: string {'average', 'min', 'max', 'dense', 'ordinal'}
// The method used to assign ranks to tied elements.
// The following methods are available:
//   - 'average': The average of the ranks that would have been assigned to
//     all the tied values is assigned to each value.
//   - 'min': The minimum of the ranks that would have been assigned to all
//     the tied values is assigned to each value.  (This is also
//     referred to as "competition" ranking.)
//   - 'max': The maximum of the ranks that would have been assigned to all
//     the tied values is assigned to each value.
//   - 'dense': Like 'min', but the rank of the next highest element is
//     assigned the rank immediately after those assigned to the tied
//     elements.
//   - 'ordinal': All values are given a distinct rank, corresponding to
//     the order that the values occur in `a`.
func RankData(a []float64, method string) []float64 {
	methods := map[string]func([]float64) []float64{
		"average": RankAverage,
		"min":     RankMin,
		"max":     RankMax,
		"dense":   RankDense,
		"ordinal": RankOrdinal,
	}
	rankFunc, found := methods[method]
	if !found {
		panic("unknown method \"" + method + "\"")
	}
	return rankFunc(a)
}
