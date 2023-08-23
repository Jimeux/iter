package secret

import (
	"slices"
	"testing"
)

func TestInts(t *testing.T) {
	ss := Ints{
		{1, true},
		{2, false},
		{3, true},
	}

	t.Run("All", func(t *testing.T) {
		got := make(Ints, 0, len(ss))
		// TODO uncomment when All is implemented
		/*for s := range ss.All() {
			got = append(got, s)
		}*/

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
		got := make(Ints, len(ss))
		// TODO uncomment when AllWithIndex is implemented
		/*for i, s := range ss.AllWithIndex() {
			got[i] = s
		}*/

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
		// TODO uncomment when All is implemented
		/*filter := iter.Filter(func(s Int) bool {
			return s.Visible
		}, ss.All())
		mapped := iter.Map(func(s Int) string {
			return strconv.Itoa(s.Val) + "!"
		}, filter)
		got := iter.Reduce(nil, func(sum []string, s string) []string {
			return append(sum, s)
		}, mapped)

		want := []string{"1!", "3!"}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}*/
	})
}
