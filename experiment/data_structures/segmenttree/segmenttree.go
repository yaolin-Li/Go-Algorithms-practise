package segmenttree

const emptyLazyNode = -1

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
