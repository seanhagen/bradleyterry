package bradleyterry

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const eps = 1e-1

func TestGetElements(t *testing.T) {
	var data = []Pair{
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 3", Loser: "Player 2"},
	}

	should := []string{"Player 1", "Player 2", "Player 3"}

	elements := getElements(data)
	sort.Strings(elements)
	if !reflect.DeepEqual(should, elements) {
		t.Errorf("Invalid elements, expected %#v got %#v", should, elements)
	}
}

func TestModel(t *testing.T) {
	rand.Seed(10)
	var data = []Pair{
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 3", Loser: "Player 2"},
	}

	should := map[string]float64{
		"Player 1": 0,
		"Player 2": 0.37,
		"Player 3": 0.63,
	}

	out := Model(data)
	for k, v := range out {
		b := should[k]
		x := math.Abs(v - b)
		if x >= eps {
			t.Errorf("Answer not close enough (0.01 decimal places) for '%v', should: %v, got %v (diff: %v)", k, b, v, x)
		}
	}
}

var benchResult = map[string]float64{}

func BenchmarkModelSmallDataset(b *testing.B) {
	var data = []Pair{
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 3", Loser: "Player 2"},
	}

	var r map[string]float64
	for n := 0; n < b.N; n++ {
		r = Model(data)
	}
	benchResult = r
}

func BenchmarkModelLargeDataset(b *testing.B) {
	var data = []Pair{
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 4", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 1", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 1", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 1", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 4"},
		{Winner: "Player 4", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 4", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 1"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 2", Loser: "Player 2"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 2"},
		{Winner: "Player 1", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 4"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 1"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 2", Loser: "Player 2"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 3", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 1", Loser: "Player 2"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 2"},
		{Winner: "Player 1", Loser: "Player 3"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 1"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 1"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 1"},
		{Winner: "Player 4", Loser: "Player 1"},
		{Winner: "Player 3", Loser: "Player 3"},
		{Winner: "Player 4", Loser: "Player 1"},
		{Winner: "Player 1", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 4"},
		{Winner: "Player 4", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 1", Loser: "Player 4"},
		{Winner: "Player 3", Loser: "Player 1"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 1", Loser: "Player 1"},
		{Winner: "Player 4", Loser: "Player 3"},
		{Winner: "Player 1", Loser: "Player 1"},
	}

	var r map[string]float64
	for n := 0; n < b.N; n++ {
		r = Model(data)
	}
	benchResult = r
}

func ExampleModel() {
	rand.Seed(10)
	var data = []Pair{
		{Winner: "Player 2", Loser: "Player 1"},
		{Winner: "Player 2", Loser: "Player 3"},
		{Winner: "Player 3", Loser: "Player 2"},
		{Winner: "Player 3", Loser: "Player 2"},
	}

	output := Model(data)
	fmt.Printf("Player 1: %.2f\n", output["Player 1"])
	fmt.Printf("Player 2: %.2f\n", output["Player 2"])
	fmt.Printf("Player 3: %.2f\n", output["Player 3"])

	// Output:
	// Player 1: 0.00
	// Player 2: 0.36
	// Player 3: 0.64
}
