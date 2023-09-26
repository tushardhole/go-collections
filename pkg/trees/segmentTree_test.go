package trees

import (
	"testing"

	. "github.com/franela/goblin"
)

type gb struct {
	G
}

// reference: https://jojozhuang.github.io/algorithm/data-structure-segment-tree/
func TestSegmentTree(t *testing.T) {
	gb := gb{G: *Goblin(t)}
	gb.Describe("Segment Tree Test", func() {
		gb.It("Find min, max, sum of full range", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 1, 4, 10, 0, 3)
		})

		gb.It("Find min, max, sum of first two elemnts", func() {
			input := []int{-1, 0, 1, 2}
			st := NewSegmentTree(input)

			gb.verify(st, -1, 0, -1, 0, 1)
		})

		gb.It("Find min, max, sum of first three elemnts", func() {
			input := []int{-1, 0, 1, 2}
			st := NewSegmentTree(input)

			gb.verify(st, -1, 1, 0, 0, 2)
		})

		gb.It("Find min, max, sum of first element", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 1, 1, 1, 0, 0)
		})

		gb.It("Find min, max, sum of last two elemnts", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 3, 4, 7, 2, 3)
		})

		gb.It("Find min, max, sum of last three elemnts", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 2, 4, 9, 1, 3)
		})

		gb.It("Find min, max, sum of last element", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 4, 4, 4, 3, 3)
		})

		gb.It("Find min, max, sum of all same elements", func() {
			input := []int{1, 1, 1, 1}
			st := NewSegmentTree(input)

			gb.verify(st, 1, 1, 4, 0, 3)
		})
	})
}

func TestSegmentUpdate(t *testing.T) {
	gb := gb{G: *Goblin(t)}
	gb.Describe("Segment Tree Update Test", func() {
		gb.It("Update first elemnt, test before & after update", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 1, 4, 10, 0, 3)

			st.Upadte(0, -1)

			gb.verify(st, -1, 4, 8, 0, 3)
		})

		gb.It("Update last elemnt, test before & after update", func() {
			input := []int{1, 2, 3, 4}
			st := NewSegmentTree(input)

			gb.verify(st, 1, 4, 10, 0, 3)

			st.Upadte(3, 11)

			gb.verify(st, 1, 11, 17, 0, 3)
		})

		gb.It("Update mid elemnt, test before & after update", func() {
			input := []int{1, 2, 3, 4, 5}
			st := NewSegmentTree(input)

			gb.verify(st, 1, 5, 15, 0, 4)

			st.Upadte(2, 1)

			gb.verify(st, 1, 5, 13, 0, 4)
		})
	})
}

func (g *gb) verify(st SegmentTree, expectedMin int, expectedMax int, expectedSum int, start int, end int) {
	min, _ := st.MinOfRange(start, end)
	max, _ := st.MaxOfRange(start, end)
	sum, _ := st.SumOfRange(start, end)

	g.Assert(min).Equal(expectedMin)
	g.Assert(max).Equal(expectedMax)
	g.Assert(sum).Equal(expectedSum)
}
