package queues

import (
	"reflect"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/tushardhole/go-collections/src/wrappers"
)

func TestPriorityQueueImpl(t *testing.T) {
	g := Goblin(t)
	g.Describe("Priority queue test", func() {
		// Passing Test
		g.It("Should add and Peek top ", func() {
			pq := NewPriorityQueue(reflect.TypeOf(ComparableInteger{}))
			pq.Add(NewComparableInteger(1))

			top := pq.Peek()
			g.Assert(top.(ComparableInteger).Value()).Equal(1)

			isEmpty := pq.IsEmpty()
			g.Assert(isEmpty).Equal(false)

			size := pq.Size()
			g.Assert(size).Equal(1)
		})

		g.It("Should add and Poll top ", func() {
			pq := NewPriorityQueue(reflect.TypeOf(ComparableInteger{}))
			pq.Add(NewComparableInteger(1))

			top := pq.Poll()
			g.Assert(top.(ComparableInteger).Value()).Equal(1)

			isEmpty := pq.IsEmpty()
			g.Assert(isEmpty).Equal(true)

			size := pq.Size()
			g.Assert(size).Equal(0)
		})

		g.It("Should add multiple return smallest on Peek", func() {
			pq := NewPriorityQueue(reflect.TypeOf(ComparableInteger{}))
			pq.Add(NewComparableInteger(1))
			pq.Add(NewComparableInteger(2))
			pq.Add(NewComparableInteger(-1))
			pq.Add(NewComparableInteger(4))
			pq.Add(NewComparableInteger(5))

			top := pq.Peek()
			g.Assert(top.(ComparableInteger).Value()).Equal(-1)
		})

		g.It("Should not allow to add invalid type", func() {

			pq := NewPriorityQueue(reflect.TypeOf(ComparableInteger{}))
			err := pq.Add(NewComparableFloat(1.0))
			g.Assert(err.Error()).Equal("extended type is not the same as the one used to create the queue, expected wrappers.ComparableInteger, got wrappers.ComparableFloat")
		})
	})
}
