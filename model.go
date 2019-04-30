package bradleyterry

import (
	"math"
	"math/rand"
)

// Pair is a struct combining the outcome of a single match, and contains the name of the
// winner and the name of the loser.
type Pair struct {
	Winner string
	Loser  string
}

const convergenceErr = 1e-05
const maxIter = 10e6

// Model takes in a dataset of Pairs that describe the winner and loser of multiple match-ups.
// The model will return a map[string]float64 that contains the preference/relevance of each
// given element.
func Model(data []Pair) map[string]float64 {
	elements := getElements(data)
	out := map[string]float64{}

	coeffVec := map[string]float64{}
	interVec := map[string]float64{}
	for _, x := range elements {
		coeffVec[x] = rand.Float64()
		interVec[x] = 0
	}

	playWins := map[string]int{}
	playPairings := map[string]map[string]int{}
	for _, pair := range data {
		win := pair.Winner
		lose := pair.Loser
		playWins[win]++

		if _, ok := playPairings[win]; !ok {
			playPairings[win] = map[string]int{}
		}
		if _, ok := playPairings[lose]; !ok {
			playPairings[lose] = map[string]int{}
		}

		playPairings[win][lose]++
		playPairings[lose][win]++
	}

	for i := 0; i < maxIter; i++ {
		for play := range coeffVec {
			ew := playWins[play]

			weightedNumPairings := 0.0
			for oe, c := range playPairings[play] {
				weightedNumPairings += float64(c) / (coeffVec[play] + coeffVec[oe])
			}

			interVec[play] = float64(ew) / weightedNumPairings
		}

		s := 0.0
		for _, v := range interVec {
			s += v
		}
		for e, coef := range interVec {
			interVec[e] = coef / s
		}

		errSum := 0.0
		for k := range coeffVec {
			v := coeffVec[k] - interVec[k]
			errSum += math.Pow(v, 2)
		}

		coeffVec = interVec

		if errSum < convergenceErr {
			return coeffVec
		}
	}
	return out
}

func getElements(input []Pair) []string {
	found := map[string]bool{}
	for _, d := range input {
		found[d.Winner] = true
		found[d.Loser] = true
	}

	elements := []string{}
	for k := range found {
		elements = append(elements, k)
	}
	return elements
}
