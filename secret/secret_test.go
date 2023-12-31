package secret

import (
	"slices"
	"strconv"
	"testing"

	"github.com/Jimeux/iter/iter"
)

func TestInts(t *testing.T) {
	ints := Ints{
		{1, true},
		{2, false},
		{3, true},
	}

	t.Run("All", func(t *testing.T) {
		got := make(Ints, 0, len(ints))
		for s := range ints.All {
			got = append(got, s)
		}

		want := Ints{
			{1, true},
			{2, false},
			{3, true},
		}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
	t.Run("AllWithIndex", func(t *testing.T) {
		got := make(Ints, len(ints))
		for i, s := range ints.AllWithIndex {
			got[i] = s
		}

		want := Ints{
			{1, true},
			{2, false},
			{3, true},
		}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
	t.Run("FilterMapReduce", func(t *testing.T) {
		filter := iter.Filter(func(s Int) bool {
			return s.Visible
		}, ints.All)
		mapped := iter.Map(func(s Int) string {
			return strconv.Itoa(s.Val) + "!"
		}, filter)
		got := iter.Reduce(nil, func(sum []string, s string) []string {
			return append(sum, s)
		}, mapped)

		want := []string{"1!", "3!"}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
}
