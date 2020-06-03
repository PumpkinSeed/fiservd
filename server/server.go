package server

type Listener interface {
	Listen() error
}

const (
	Port = ":2345"
)

func Listen() error {
	return NewSI().Listen()
}
