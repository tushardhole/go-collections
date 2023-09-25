package behaviours

type Comparable interface {
	CompareTo(other Comparable) int
}
