package container

// Foreach function iterates every node in container and does f(args).
func Foreach(c Container, f func(n Node, idx int)) {
	if c == nil || f == nil {
		return
	}
	for i := 0; i < c.Len(); i++ {
		n := c.At(i)
		f(n, i)
	}
}
