package server

import (
	"bufio"
	"fmt"
	"github.com/PumpkinSeed/fiservd/server/handler"
	"net"
)

// si - Single but interrupt previous connection
//
// This means when ONE client connects it has direct connection to the server
// until another connection not comes in. When the new connection comes in,
// the previous connection get closed and the new connection become the only one valid connection.

type SI struct {
	conns []conn
}

type conn struct {
	c net.Conn

	// use channel becase: possible data race can occur, when the newConn get closed sooner
	closeCh chan struct{}
}

func NewSI() Listener {
	return &SI{
		conns: []conn{},
	}
}

func (s *SI) Listen() error {
	l, err := net.Listen("tcp", Port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		s.conns = append(s.conns, newConn(c))
		s.closePrev()
		go handleConn(s.conns[s.last()])
	}

	return nil
}

// close the previous connection
func (s *SI) closePrev() {
	if len(s.conns) > 1 && s.conns[s.last()-1].c != nil {
		s.conns[s.last()-1].closeCh <- struct{}{}
		s.conns[s.last()-1].c.Close()
	}
}

func (s *SI) last() int {
	return len(s.conns) - 1
}

func newConn(c net.Conn) conn {
	return conn{
		c:       c,
		closeCh: make(chan struct{}),
	}
}

func handleConn(c conn) {
	for {
		select {
		case <-c.closeCh:
			fmt.Println("Connection closed")
			return
		default:
			netData, err := bufio.NewReader(c.c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			response := handler.Handle(string(netData))
			c.c.Write([]byte(response))
		}
	}
}
