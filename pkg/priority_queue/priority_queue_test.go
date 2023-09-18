package priority_queue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPQ(t *testing.T) {
	items := map[string]int64{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}

	pq := make(priorityQueue[string], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = NewItem(value, priority, i, SetProperty[string]("internal", 123))
		//pq[i] = Item[string]{Value: value, Priority: priority, Index: i}
		i++
	}
	heap.Init(&pq)

	item := &Item[string]{Value: "orange", Priority: 1}
	heap.Push(&pq, item)
	pq.update(item, "orange", 5)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(Item[string])
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
}
