package api

func ExampleAddMessage() {
	room := Room{Messages: []*Message{}}

	room.AddMessage(Message{"Hello, World!"})
	// Output: api.Message{Value:"Hello, World!"}
}
