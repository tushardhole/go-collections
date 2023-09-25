package trees

import (
	"testing"

	. "github.com/franela/goblin"
)

// reference: https://jojozhuang.github.io/algorithm/data-structure-segment-tree/
func TestSegmentTree(t *testing.T) {
	g := Goblin(t)
	g.Describe("Segment Tree Test", func() {
		g.It("Find min of full range", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(0, 3)

			g.Assert(minOfAll).Equal(1)
		})

		g.It("Find min of first two elemnts", func() {
			input := []int{-1, 0, 1, 2}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(0, 1)

			g.Assert(minOfAll).Equal(-1)
		})

		g.It("Find min of first three elemnts", func() {
			input := []int{-1, 0, 1, 2}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(0, 2)

			g.Assert(minOfAll).Equal(-1)
		})

		g.It("Find min of first element", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(0, 0)

			g.Assert(minOfAll).Equal(1)
		})

		g.It("Find min of last two elemnts", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(2, 3)

			g.Assert(minOfAll).Equal(3)
		})

		g.It("Find min of last three elemnts", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(1, 3)

			g.Assert(minOfAll).Equal(2)
		})

		g.It("Find min of last element", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(3, 3)

			g.Assert(minOfAll).Equal(4)
		})

		g.It("Find min of all same elements", func() {
			input := []int{1, 1, 1, 1}
			st := NewSegmentTree(input)

			minOfAll, _ := st.minOfRange(0, 3)

			g.Assert(minOfAll).Equal(1)
		})
	})
}
