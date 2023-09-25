package trees

import (
	"testing"

	. "github.com/franela/goblin"
)

// reference: https://jojozhuang.github.io/algorithm/data-structure-segment-tree/
func TestSegmentTree(t *testing.T) {
	g := Goblin(t)
	g.Describe("Segment Tree Test", func() {
		g.It("Find min, max, sum of full range", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(0, 3)

			g.Assert(minOfAll).Equal(1)
		})

		g.It("Find min, max, sum of first two elemnts", func() {
			input := []int{-1, 0, 1, 2}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(0, 1)

			g.Assert(minOfAll).Equal(-1)
		})

		g.It("Find min, max, sum of first three elemnts", func() {
			input := []int{-1, 0, 1, 2}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(0, 2)

			g.Assert(minOfAll).Equal(-1)
		})

		g.It("Find min, max, sum of first element", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(0, 0)

			g.Assert(minOfAll).Equal(1)
		})

		g.It("Find min, max, sum of last two elemnts", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(2, 3)

			g.Assert(minOfAll).Equal(3)
		})

		g.It("Find min, max, sum of last three elemnts", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(1, 3)

			g.Assert(minOfAll).Equal(2)
		})

		g.It("Find min, max, sum of last element", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(3, 3)

			g.Assert(minOfAll).Equal(4)
		})

		g.It("Find min, max, sum of all same elements", func() {
			input := []int{1, 1, 1, 1}
			st := NewSegmentTree(input)

			minOfAll, _ := st.MinOfRange(0, 3)

			g.Assert(minOfAll).Equal(1)
		})
	})
}
