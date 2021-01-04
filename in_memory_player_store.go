package main

import "sync"

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{sync.Mutex{}, map[string]int{}}
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.store[name]

}

// RecordWin will record a player's win.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

// GetLeague will return the league data from the store.
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
