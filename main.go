package main

import (
	"fmt"
	"os"

	"github.com/PumpkinSeed/fiservd/client"
	"github.com/PumpkinSeed/fiservd/server"
)

func main() {
	fmt.Println(os.Args[1])
	if len(os.Args) > 1 && os.Args[1] == "server" {
		if err := server.Listen(); err != nil {
			panic(err)
		}
	} else {
		if err := client.Connect("localhost", server.Port); err != nil {
			fmt.Println(err)
		}
	}
}
