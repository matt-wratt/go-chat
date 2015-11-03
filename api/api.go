package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

/*

GET  /api/rooms
POST /api/rooms
PUT  /api/rooms/:id
GET  /api/rooms/:id
POST /api/rooms/:id
GET  /client

*/

var store = Store{Rooms: []*Room{}}

// Route doc
type Route struct {
	method  string
	path    *regexp.Regexp
	handler http.HandlerFunc
}

// NewRoute doc
func NewRoute(method, path string, handler http.HandlerFunc) Route {
	return Route{
		method:  method,
		path:    regexp.MustCompile(path),
		handler: handler,
	}
}

func (route Route) match(method, path string) bool {
	if method == route.method && route.path.MatchString(path) {
		return true
	}
	return false
}

// API doc
type API struct {
	routes []Route
}

// NewAPI doc
func NewAPI() API {
	return API{routes: []Route{
		NewRoute("GET", "^/api/rooms/?$", listRooms),
		NewRoute("POST", "^/api/rooms/?$", createRoom),
		NewRoute("GET", "^/api/rooms/.*/?$", showRoom),
		NewRoute("POST", "^/api/rooms/.*/?$", addRoomMessage),
		NewRoute("PUT", "^/api/rooms/.*/?$", updateRoom),
	}}
}

// APIHandleFunc routes requests to api resources
func (api API) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	found := false
	method := r.Method
	path := r.URL.Path

	log.Printf("[%s] %s", method, path)

	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT")

	if method == "OPTIONS" {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	for _, route := range api.routes {
		if route.match(method, path) {
			route.handler(w, r)
			found = true
			break
		}
	}

	if found == false {
		http.NotFound(w, r)
	}

}

func listRooms(w http.ResponseWriter, r *http.Request) {
	if json, err := json.Marshal(store.Rooms); err == nil {
		w.Write(json)
	} else {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	opt := RoomOptions{}
	json.NewDecoder(r.Body).Decode(&opt)
	store.AddRoom(opt)
	listRooms(w, r)
}

func showRoom(w http.ResponseWriter, r *http.Request) {
	if room, err := roomFromRequest(r); err == nil {
		if json, err := json.Marshal(room); err == nil {
			w.Write(json)
		} else {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
	} else {
		http.NotFound(w, r)
	}
}

func updateRoom(w http.ResponseWriter, r *http.Request) {
	if room, err := roomFromRequest(r); err == nil {
		if err := json.NewDecoder(r.Body).Decode(room); err == nil {
			showRoom(w, r)
		} else {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
	} else {
		http.NotFound(w, r)
	}
}

func addRoomMessage(w http.ResponseWriter, r *http.Request) {
	if room, err := roomFromRequest(r); err == nil {
		var message Message
		if err := json.NewDecoder(r.Body).Decode(&message); err == nil {
			room.AddMessage(message)
			showRoom(w, r)
		} else {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
	} else {
		http.NotFound(w, r)
	}
}

func roomFromRequest(r *http.Request) (room *Room, err error) {
	path := strings.TrimLeft(r.URL.Path, "/api/rooms/")
	if room, err := store.Room(path); err == nil {
		return room, nil
	}
	return room, fmt.Errorf("Room not found '%v'", path)
}
