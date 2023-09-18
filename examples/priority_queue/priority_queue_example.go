package main

import (
	"fmt"
	"github.com/jaylee630.go-structs-kit/pkg/priority_queue"
	"sync"
)

func main() {
	pq := priority_queue.New[string]()

	// 添加元素
	opts1 := priority_queue.SetProperty[string]("color", "red")
	item1 := priority_queue.NewItem("low", 10, 1, opts1)

	opts2 := priority_queue.SetProperty[string]("color", "blue")
	item2 := priority_queue.NewItem("medium", 5, 2, opts2)

	opts3 := priority_queue.SetProperty[string]("color", "black")
	item3 := priority_queue.NewItem("high", 1, 3, opts3)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		pq.Push(item1)
		pq.Push(item2)
		pq.Push(item3)
	}()
	wg.Wait()

	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("Value: %v, Priority: %d\n", item.Value, item.Priority)

		// 打印额外的属性
		if color, ok := item.Properties("color").(string); ok {
			fmt.Printf("Color: %s\n", color)
		}
	}
}
