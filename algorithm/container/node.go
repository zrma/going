package container

// Container interface implements iteration.
type Container interface {
	Len() int
	At(idx int) Node
}

// Node interface can be contained in container.
type Node interface {
	isNode()
}
