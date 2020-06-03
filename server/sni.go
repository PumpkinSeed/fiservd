package server

import (
	"bufio"
	"github.com/PumpkinSeed/fiservd/server/handler"
	"net"
)

// sni - Single non interrupt connection
//
// This means when ONE client connects it can have a direct non-interrupted conntection.
// When a second client connects it won't interrupt the first's connection, so it's packages
// won't served until the first connection not closed.

type SNI struct{}

func NewSNI() Listener {
	return SNI{}
}

func (SNI) Listen() error {
	l, err := net.Listen("tcp", Port)
	if err != nil {
		return err
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		return err
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return err
		}

		response := handler.Handle(string(netData))
		c.Write([]byte(response))
	}

	return nil
}
