package core

import (
	"github.com/gofxh/blog/lib/log"
	"github.com/lunny/tango"
	"net"
	"net/http"
	"time"
)

type Server struct {
	address string
	*tango.Tango
	ln      net.Listener
	isClose bool
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
		Tango:   tango.Classic(),
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Server|Start|%s", err.Error())
	}
	s.ln = ln

	httpServer := &http.Server{Addr: s.address, Handler: s.Tango}
	Wrap("Server|Listen", func() {
		if err = httpServer.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
			if s.isClose {
				return
			}
			log.Fatal("Server|Start|%s", err.Error())
		}
	})
}

func (s *Server) Stop() {
	s.isClose = true
	if s.ln != nil {
		s.ln.Close()
	}
}

// copy from net/http/server.go
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
