package main

import (
	"fmt"
	"context"
)

//迭代器实现
// 使用数组做底层存储
type arrBag struct {
	data []interface{}
	size int
}

// 使用链表做底层存储
type node struct {
	data interface{}
	next *node
}
type linkBag struct {
	node *node
	size int
}

//闭包实现
// For Array underlying implementation
func (a *arrBag) Iterator() (func() (interface{}, bool), bool) {
	index := 0
	var item interface{}
	return func() (interface{}, bool) {
		if index < a.size {
			item = a.data[index]
			index++
		} else {
			item = nil
		}
		return item, index < a.size
	}, index < a.size
}

// For Link underlying implementation
func (b *linkBag) Iterator() (func() (interface{}, bool), bool) {
	item := b.node
	return func() (interface{}, bool) {
		current := item
		if current != nil {
			item = current.next
		}
		return current, item != nil
	}, item != nil && item.next != nil
}

func useMain() {
	a := arrBag{}
	//init
	it, hasNext := a.Iterator()
	var item interface{}
	for hasNext {
		item, hasNext = it()
		fmt.Printf("%v\n", item)
	}

	//chan
	ab := arrBag{}
	var matched bool
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // 如果需要遍历所有，可以把 cancel() 放到 defer 处
	for item := range ab.Chan(ctx) {
		if item == matched {
			cancel() // 保证中断遍历，资源回收
			return
		}
	}
}

//chan定义
// For Array underlying implementation
func (b *arrBag) Chan(ctx context.Context) <-chan interface{} {
	buf := make(chan interface{})
	go func() {
		defer close(buf)
		for _, item := range b.data {
			select {
			case <-ctx.Done():
				return
			default:
				buf <- item
			}
		}
	}()
	return buf
}

// For Link underlying implementation
func (b *linkBag) Chan(ctx context.Context) <-chan interface{} {
	buf := make(chan interface{})
	node := b.node
	go func() {
		defer close(buf)
		for node != nil {
			select {
			case <-ctx.Done():
				return
			default:
				buf <- node.data
				node = node.next
			}
		}
	}()
	return buf
}
