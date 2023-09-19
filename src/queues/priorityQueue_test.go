package collections

import (
	. "github.com/franela/goblin"
	collections "go-collections/src/common"
	"testing"
)

type integer struct {
	value int
}

func (i integer) CompareTo(other collections.Comparable) int {
	return i.value - other.(integer).value
}

func TestPriorityQueueImpl(t *testing.T) {
	g := Goblin(t)
	g.Describe("Priority queue test", func() {
		// Passing Test
		g.It("Should add and Peek top ", func() {
			pq := NewPriorityQueue()
			pq.Add(integer{value: 1})

			top := pq.Peek()
			g.Assert(top.(integer).value).Equal(1)

			isEmpty := pq.IsEmpty()
			g.Assert(isEmpty).Equal(false)

			size := pq.Size()
			g.Assert(size).Equal(1)
		})

		g.It("Should add and Poll top ", func() {
			pq := NewPriorityQueue()
			pq.Add(integer{value: 1})

			top := pq.Poll()
			g.Assert(top.(integer).value).Equal(1)

			isEmpty := pq.IsEmpty()
			g.Assert(isEmpty).Equal(true)

			size := pq.Size()
			g.Assert(size).Equal(0)
		})

		g.It("Should add multiple return smallest on Peek", func() {
			pq := NewPriorityQueue()
			pq.Add(integer{value: 3})
			pq.Add(integer{value: 1})
			pq.Add(integer{value: 8})
			pq.Add(integer{value: 0})
			pq.Add(integer{value: -1})

			top := pq.Peek()
			g.Assert(top.(integer).value).Equal(-1)
		})
	})
}
