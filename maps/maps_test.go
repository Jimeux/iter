package maps

import (
	"slices"
	"testing"
)

func TestIterators(t *testing.T) {
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	t.Run("Keys", func(t *testing.T) {
		got := make([]string, 0, len(m))
		// TODO 2023/08/24 @Jimeux uncomment when Keys is implemented
		/*for k := range Keys(m) {
			got = append(got, k)
		}*/
		slices.Sort(got)

		want := []string{"A", "B", "C"}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
	t.Run("Values", func(t *testing.T) {
		got := make([]int, 0, len(m))
		// TODO 2023/08/24 @Jimeux uncomment when Values is implemented
		/*for v := range Values(m) {
			got = append(got, v)
		}*/
		slices.Sort(got)

		want := []int{1, 2, 3}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
}
