package main

import (
	"fmt"
	"github.com/jaylee630.go-structs-kit/pkg/priority_queue"
)

func main() {
	pq := priority_queue.New[int]()

	// 添加元素
	opts1 := priority_queue.SetProperty[int]("color", "red")
	item1 := priority_queue.NewItem(1, 10, 1, opts1)
	pq.HeapPush(item1)

	opts2 := priority_queue.SetProperty[int]("color", "blue")
	item2 := priority_queue.NewItem(2, 5, 2, opts2)
	pq.HeapPush(item2)

	// 弹出元素并打印
	for len(pq) > 0 {
		item := pq.HeapPop()
		fmt.Printf("Value: %v, Priority: %d\n", item.Value, item.Priority)

		// 打印额外的属性
		if color, ok := item.Properties("color").(string); ok {
			fmt.Printf("Color: %s\n", color)
		}
	}
}
