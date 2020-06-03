package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Connect(host, port string) error {
	c, err := net.Dial("tcp", host+port)
	if err != nil {
		return err
	}

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Print("->: " + message)
	}

	return nil
}
