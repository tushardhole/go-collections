package collections

import (
	"errors"
	"fmt"
	. "go-collections/src/common/behaviours"
	. "go-collections/src/common/types"
	"reflect"
)

type PriorityQueueImpl struct {
	heap     []Comparable
	itemType reflect.Type
}

func (pq *PriorityQueueImpl) Size() int {
	return len(pq.heap)
}

func NewPriorityQueue(itemType reflect.Type) PriorityQueue {
	return &PriorityQueueImpl{itemType: itemType}
}

func (pq *PriorityQueueImpl) Add(item Comparable) error {
	validationErr := pq.validate(item)
	if validationErr != nil {
		return validationErr
	}

	pq.heap = append(pq.heap, item)
	pq.heapifyUp(len(pq.heap) - 1)

	return nil
}

func (pq *PriorityQueueImpl) validate(item Comparable) error {
	if pq.itemType != reflect.TypeOf(item) {
		errorsMsg := fmt.Sprintf("extended type is not the same as the one used to create the queue, expected %s, got %s", pq.itemType, reflect.TypeOf(item))
		return errors.New(errorsMsg)
	}
	return nil
}

func (pq *PriorityQueueImpl) heapifyUp(i int) {
	parent := (i - 1) / 2
	if parent < 0 {
		return
	}
	if pq.heap[parent].CompareTo(pq.heap[i]) > 0 {
		pq.heap[parent], pq.heap[i] = pq.heap[i], pq.heap[parent]
		pq.heapifyUp(parent)
	}
}

func (pq *PriorityQueueImpl) Poll() Comparable {
	if pq.IsEmpty() {
		return nil
	}
	item := pq.heap[0]
	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	pq.heapifyDown(0)
	return item
}

func (pq *PriorityQueueImpl) heapifyDown(i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < len(pq.heap) && pq.heap[left].CompareTo(pq.heap[smallest]) < 0 {
		smallest = left
	}
	if right < len(pq.heap) && pq.heap[right].CompareTo(pq.heap[smallest]) < 0 {
		smallest = right
	}
	if smallest != i {
		pq.heap[smallest], pq.heap[i] = pq.heap[i], pq.heap[smallest]
		pq.heapifyDown(smallest)
	}
}

func (pq *PriorityQueueImpl) Peek() Comparable {
	if pq.IsEmpty() {
		return nil
	}
	return pq.heap[0]
}

func (pq *PriorityQueueImpl) IsEmpty() bool {
	return len(pq.heap) == 0
}
