package priority_queue

import (
	"container/heap"
	"sync"
	"time"
)

type Item[T any] struct {
	Value    T
	Priority int64
	Index    int

	insertTime int64

	properties map[string]any
}

// NewItem 创建新的Item
func NewItem[T any](value T, priority int64, index int, opts ...Option[T]) *Item[T] {
	item := &Item[T]{Value: value, Priority: priority, Index: index, insertTime: time.Now().Unix()}
	for _, opt := range opts {
		opt(item)
	}
	return item
}

func (i *Item[T]) Properties(key string) any {
	val, _ := i.properties[key]
	return val
}

type priorityQueue[T any] []*Item[T]

func (pq *priorityQueue[T]) Len() int {
	return len(*pq)
}

// Less Priority小优先, 早到优先
func (pq *priorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Priority+(*pq)[i].insertTime < (*pq)[j].Priority+(*pq)[j].insertTime
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].Index, (*pq)[j].Index = i, j
}

func (pq *priorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = &Item[T]{Index: -1}
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue[T]) update(item *Item[T], value T, priority int64) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

type PriorityQueue[T any] struct {
	mu sync.Mutex
	pq priorityQueue[T]
}

func New[T any]() *PriorityQueue[T] {
	pq := priorityQueue[T]{}
	heap.Init(&pq)
	return &PriorityQueue[T]{
		pq: pq,
	}
}

func (pq *PriorityQueue[T]) Push(item *Item[T]) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	heap.Push(&pq.pq, item)
}

// Pop 目前不支持并发调用
func (pq *PriorityQueue[T]) Pop() *Item[T] {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return heap.Pop(&pq.pq).(*Item[T])
}

func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int64) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.pq.update(item, value, priority)
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pq.Len()
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
