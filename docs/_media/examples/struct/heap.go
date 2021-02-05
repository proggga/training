package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	hp := &heap{}
	hp.init(10)
	hp.add(10)
	hp.add(9)
	hp.add(8)
	hp.add(7)
	hp.add(6)
	hp.add(11)
	hp.add(5)
	hp.pop()
	hp.add(1)
	hp.pop()
	hp.add(21)
	hp.add(1)
	hp.add(31)
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	fmt.Printf("%+v\n", hp)
	sm := []int32{1,6,2,3,5,67,7,23,4,9}
	hp.setNewHeap(sm)
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	hp.pop()
	fmt.Printf("%+v\n", hp)
}

type heap struct {
	storage []int32
	cursor  int
}

func (h *heap) add(value int32) {
	h.cursor += 1
	h.storage[h.cursor] = value
	h.heapify(h.cursor)
	fmt.Printf("add %d\n", value)
}

func (h *heap) heapify(index int) {
	if index > 0 {
		parent := h.getParent(index)
		if h.storage[parent] > h.storage[index] {
			h.swap(parent, index)
			h.heapify(parent)
		}
	}
}

func (h *heap) swap(first, second int) {
	h.storage[first], h.storage[second] = h.storage[second], h.storage[first]
}

func (h *heap) getLeftChild(index int) int {
	return index*2 + 1
}

func (h *heap) getRightChild(index int) int {
	return index*2 + 2
}

func (h *heap) getParent(index int) int {
	return index / 2
}

func (h *heap) init(size int) {
	h.storage = make([]int32, size)
	h.cursor = -1
}

func (h *heap) pop() int32 {
	if h.cursor == -1 {
		return -1
	}
	value := h.storage[0]
	h.swap(0, h.cursor)
	h.storage[h.cursor] = 0
	h.cursor -= 1
	h.minHeapify(0)
	fmt.Printf("pop %d\n", value)
	return value
}

func (h *heap) minHeapify(index int) {
	left_i := h.getLeftChild(index)
	right_i := h.getRightChild(index)


	smallest := index
	if left_i <= h.cursor &&  h.storage[left_i] < h.storage[index] {
		smallest = left_i
	}
	if right_i <= h.cursor && h.storage[right_i] < h.storage[smallest] {
		smallest = right_i
	}
	if smallest != index {
		h.swap(smallest, index)
		h.minHeapify(smallest)
	}
}

func (h *heap) setNewHeap(values []int32) {
	h.init(len(values))
	copy(h.storage, values)
	h.cursor = len(values) - 1
	fmt.Printf("cur %d - %+v\n", h.cursor, h.storage)
	h.rebuildHeap()
}

func (h *heap) rebuildHeap() {
	current_cursor := h.cursor
	for i := 0; i < current_cursor; i++ {
		h.heapify(i)
	}
}
