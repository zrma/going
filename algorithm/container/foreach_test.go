package container

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("컨테이너 동작 검증", func() {
	Context("Foreach", func() {
		It("빈 컨테이너를 인자로 전달해도 잘 동작한다.", func(done Done) {
			defer GinkgoRecover()
			Foreach(nil, nil)
			close(done)
		})

		It("Iterate 동작을 잘 수행한다.", func(done Done) {
			defer GinkgoRecover()

			const size = 100
			var nodes []node
			for i := 0; i < size; i++ {
				nodes = append(nodes, node(i))
			}
			c := containerImpl{nodes}
			Foreach(&c, func(n Node, idx int) {
				Expect(n).Should(Equal(nodes[idx]))
			})
			close(done)
		})
	})
})

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
