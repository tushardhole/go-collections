package collections

type Comparable interface {
	CompareTo(other Comparable) int
}
