package main

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{map[string]int{}}
}

type InMemoryUserStore struct {
	store map[string]int
}

func (i *InMemoryUserStore) GetUserId(id string) int {
	return i.store[id]
}

func (i *InMemoryUserStore) RecordUser(id string) {
	i.store[id]++
}