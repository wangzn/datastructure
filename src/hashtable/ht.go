package ht

import (
	"errors"
	"fmt"
	"math"

	"github.com/wangzn/datastructure/src/linked-list"
)

type HashTable struct {
	table    map[int]*list.List
	size     int
	capacity int
}

type item struct {
	key   string
	value interface{}
}

func New(capacity int) *HashTable {
	t := make(map[int]*list.List, capacity)
	return &HashTable{
		table:    t,
		size:     0,
		capacity: capacity,
	}
}

func (i *item) Print() {
	fmt.Println("----------")
	fmt.Println(i.key, ":", i.value)
	fmt.Println("----------")
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) Get(key string) (interface{}, error) {
	pos, err := h.getPosition(key)
	if err != nil {
		return nil, err
	}
	if _, ok := h.table[pos]; !ok {
		return nil, errors.New("Key position is nil")
	}
	v, err := h.find(pos, key)
	if err != nil || v == nil {
		return nil, err
	}
	return v.value, nil
}

func (h *HashTable) Set(key string, value interface{}) error {
	pos, err := h.getPosition(key)
	if err != nil {
		return err
	}
	data := &item{
		key:   key,
		value: value,
	}
	if _, ok := h.table[pos]; !ok {
		h.table[pos] = list.NewList()
	}
	p, err := h.find(pos, key)
	if err != nil || p == nil {
		h.table[pos].Append(data)
	} else {
		p.value = value
	}
	h.size++
	return nil
}

func (h *HashTable) Del(key string) error {
	if h == nil || h.size == 0 {
		return errors.New("Hash table is empty")
	}
	pos, err := h.getPosition(key)
	if err != nil {
		return err
	}
	p, err := h.find(pos, key)
	if err != nil || p == nil {
		return errors.New("Key doesnot exist")
	}
	r := h.table[pos].Remove(p)
	if r <= 0 {
		return errors.New("Remove value fail in list")
	}
	h.size = h.size - r
	return nil
}

func (h *HashTable) Walk(f func(i *item)) {
	for _, v := range h.table {
		if v != nil {
			v.Walk(func(n *list.Node) {
				f(n.Value().(*item))
			})
		}
	}
}

func (h *HashTable) Print() {
	for k, v := range h.table {
		if v != nil {
			fmt.Println("+Slot", k, "+")
			v.Walk(func(n *list.Node) {
				i := n.Value().(*item)
				i.Print()
			})
		}
	}
}

func (h *HashTable) find(pos int, key string) (*item, error) {
	if _, ok := h.table[pos]; !ok {
		return nil, errors.New("pos does not exist")
	}
	var val *item
	h.table[pos].Walk(func(n *list.Node) {
		if n.Value().(*item).key == key {
			val = n.Value().(*item)
		}
	})
	if val == nil {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (h *HashTable) getPosition(key string) (int, error) {
	if h == nil || h.capacity == 0 {
		return 0, errors.New("Hash table is empty")
	}
	return getHash(key) % h.capacity, nil
}

func getHash(key string) int {
	hash := int32(0)
	for i := 0; i < len(key); i++ {
		hash = int32(hash<<5-hash) + int32(key[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}
