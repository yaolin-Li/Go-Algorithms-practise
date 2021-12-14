package segmenttree

const emptyLazyNode = -1

import (
	"github.com/Go-Algorithms-practise/math/min"
)
type SegmentTree struct {
	Array       []int
	SegmentTree []int
	LazyTree    []int
}

func (s *SegmentTree) Propagate(node int, leftNode int, rightNode int) {
	if s.LazyTree[node] != emptyLazyNode {

		s.SegmentTree[node] += (rightNode - leftNode + 1) * s.LazyTree[node]

		if leftNode == rightNode {

			s.Array[leftNode] = s.LazyTree[node]
		} else {
			s.LazyTree[2*node] = s.LazyTree[node]
			s.LazyTree[2*node+1] = s.LazyTree[node]
		}
		s.LazyTree[node] = emptyLazyNode
	}
}

func (s *SegmentTree) Query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) int {
	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		return 0
	}

	s.Propagate(node, leftNode, rightNode)

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {
		return s.SegmentTree[node]
	}

	mid := (leftNode + rightNode) / 2

	leftNodeSum := s.Query(2 * node, leftNode, mid, firstIndex, min)
}