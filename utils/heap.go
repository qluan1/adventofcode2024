package utils

type Comparable interface {
	IsLess(x interface{}) bool
}

type HeapItem struct {
	item Comparable
	index int
}

type MyHeap []*HeapItem

func (pq MyHeap) Len() int { return len(pq) }

func (pq MyHeap) Less(i, j int) bool {
	return pq[i].item.IsLess(pq[j].item)
}

func (pq MyHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MyHeap) Push(x interface{}) {
	n := len(*pq)
	comp, ok := x.(Comparable)
	if (!ok) {
		panic("Push() given non-comparable argument")
	}
	item := &HeapItem{
		item: comp,
		index: n,
	}
	*pq = append(*pq, item)
}

func (pq *MyHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	top := old[n-1]
	old[n-1] = nil
	top.index = -1
	*pq = old[0 : n-1]
	return top.item
}

func (pq MyHeap) Peek() Comparable {
	return pq[0].item
}
