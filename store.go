package main

import "fmt"

// Store doc
type Store struct {
	Rooms []*Room
}

// AddRoom doc
func (store *Store) AddRoom(opt RoomOptions) (r Room, err error) {
	r = NewRoom(opt)
	store.Rooms = append(store.Rooms, &r)
	return r, err
}

// Room doc
func (store Store) Room(path string) (r *Room, err error) {
	for _, r := range store.Rooms {
		if r.Path == path {
			return r, nil
		}
	}
	return r, fmt.Errorf("room not found at %v", path)
}
