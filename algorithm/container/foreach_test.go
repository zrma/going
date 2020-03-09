package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForeach(t *testing.T) {
	t.Run("컨테이너 동작 검증", func(t *testing.T) {
		t.Parallel()

		t.Run("빈 컨테이너를 인자로 전달해도 잘 동작한다.", func(t *testing.T) {
			Foreach(nil, nil)
		})

		t.Run("Iterate 동작을 잘 수행한다.", func(t *testing.T) {
			const size = 100
			var nodes []node
			for i := 0; i < size; i++ {
				nodes = append(nodes, node(i))
			}
			c := containerImpl{nodes}
			Foreach(&c, func(n Node, idx int) {
				assert.Equal(t, n, nodes[idx])
			})
		})
	})
}

type node int

func (n node) isNode() {

}

type containerImpl struct {
	data []node
}

func (c containerImpl) Len() int {
	return len(c.data)
}

func (c containerImpl) At(idx int) Node {
	if c.Len() <= idx {
		return nil
	}
	return c.data[idx]
}
