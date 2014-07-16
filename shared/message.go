package shared

// The message structure used to communicate client/server.
// Any change to this will make versions incompatible, do so with care!
type Message struct {
	Id   int64
	Name string
	Type string
	X, Y int
}
