package bridge

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	Port = ":3000"
)

type Wrapper struct {
	Data string `json:"data"`
}

type server struct {
	conn   net.Conn
	router *chi.Mux
}

func NewServer(host, port string) (server, error) {
	c, err := net.Dial("tcp", host+port)
	if err != nil {
		return server{}, err
	}

	return server{
		conn: c,
	}, nil
}

func (s server) Serve() error {
	s.router = chi.NewRouter()
	s.setMiddlewares()


	s.router.Get("/", s.handler)
	return http.ListenAndServe(Port, s.router)
}

func (s server) setMiddlewares() {
	s.router.Use(middleware.Logger)
	// Stop processing after 30 seconds.
	s.router.Use(middleware.Timeout(30 * time.Second))
	// Only one request will be processed at a time.
	s.router.Use(middleware.Throttle(1))
}

func (s server) handler(w http.ResponseWriter, r *http.Request) {
	log.Print("---- Incoming request, start to handle it")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	var wrapper Wrapper
	err = json.Unmarshal(body, &wrapper)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(s.conn, wrapper.Data+"\n")

	message, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := json.Marshal(Wrapper{message})
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write(resp)
}
