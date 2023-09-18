package priority_queue

import (
	"container/heap"
)

type Item[T any] struct {
	Value    T
	Priority int
	Index    int

	properties map[string]any
}

// NewItem 创建新的Item
func NewItem[T any](value T, priority int, index int, opts ...Option[T]) *Item[T] {
	item := &Item[T]{Value: value, Priority: priority, Index: index}
	for _, opt := range opts {
		opt(item)
	}
	return item
}

type PriorityQueue[T any] []*Item[T]

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Priority > (*pq)[j].Priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].Index, (*pq)[j].Index = i, j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = &Item[T]{Index: -1}
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

type Option[T any] func(*Item[T])

// SetProperty 设置单个属性
func SetProperty[T any](key string, value any) Option[T] {
	return func(item *Item[T]) {
		if item.properties == nil {
			item.properties = make(map[string]any)
		}
		item.properties[key] = value
	}
}

// SetProperties 设置多个属性
func SetProperties[T any](props map[string]any) Option[T] {
	return func(item *Item[T]) {
		if item.properties == nil {
			item.properties = make(map[string]any)
		}
		for k, v := range props {
			item.properties[k] = v
		}
	}
}
