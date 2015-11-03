package api

import (
	"fmt"
	"regexp"
	"strings"
)

// Room doc
type Room struct {
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Messages []*Message `json:"messages"`
}

// RoomOptions doc
type RoomOptions struct {
	Name string `json:"name"`
}

// NewRoom doc
func NewRoom(opt RoomOptions) Room {
	path := regexp.MustCompile("[^a-z0-9]+").ReplaceAllString(strings.ToLower(opt.Name), "")
	return Room{Name: opt.Name, Path: path, Messages: []*Message{}}
}

// AddMessage doc
func (room *Room) AddMessage(message Message) {
	fmt.Printf("%#v", message)
	room.Messages = append(room.Messages, &message)
}
