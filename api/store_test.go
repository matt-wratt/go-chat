package api

import "testing"

func TestAddRoom(t *testing.T) {
	store := Store{}

	store.AddRoom(RoomOptions{"Test Room"})

	if len(store.Rooms) != 1 || store.Rooms[0].Name != "Test Room" {
		t.Errorf("Room not added")
	}
}

func TestRoom(t *testing.T) {
	store := Store{
		[]*Room{
			{
				Name: "Foobar Room",
				Path: "foobar-room",
			},
			{
				Name: "Test Room",
				Path: "test-room",
			},
		},
	}

	rm, err := store.Room("test-room")
	if err != nil || rm.Name != "Test Room" {
		t.Errorf("Failed to get room")
	}

	rm, err = store.Room("wrong-path")
	if err == nil || err.Error() != "room not found at wrong-path" {
		t.Errorf("WTF")
	}
}
