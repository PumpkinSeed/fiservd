package main

import (
	"fmt"
	"github.com/PumpkinSeed/fiservd/bridge"
	"github.com/PumpkinSeed/fiservd/client"
	"github.com/PumpkinSeed/fiservd/server"
	"time"
)

//func main() {
//	fmt.Println(os.Args[1])
//	if len(os.Args) > 1 && os.Args[1] == "server" {
//		if err := server.Listen(); err != nil {
//			panic(err)
//		}
//	} else {
//		if err := client.Connect("localhost", server.Port); err != nil {
//			fmt.Println(err)
//		}
//	}
//}

func main() {
	go func() {
		if err := server.Listen(); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Server started")

	srv, err := bridge.NewServer("localhost", server.Port)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := srv.Serve(); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Bridge started")

	if err := client.Load("http://localhost", bridge.Port); err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(20*time.Second)
}
