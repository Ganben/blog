package core

import (
	"github.com/gofxh/blog/app/log"
	"github.com/lunny/tango"
	"net"
	"net/http"
	"time"
)

// http server
type Server struct {
	address string
	*tango.Tango
	ln      net.Listener // use listener to try close
	isClose bool
}

// new http server with address
func NewServer(address string) *Server {
	s := &Server{
		address: address,
	}
	// use custom tango, not classic
	s.Tango = tango.NewWithLog(log.Get().ToTangoLogger(), []tango.Handler{
		tango.Logging(),
		tango.Recovery(true),
		tango.Return(),
		tango.Param(),
		tango.Contexts(),
	}...)
	return s
}

// start http server
func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Server|Start|%s", err.Error())
	}
	s.ln = ln

	// init http server
	httpServer := &http.Server{Addr: s.address, Handler: s.Tango}

	// use global wrapper to listen server
	Wrap("Server|Listen", func() {
		if err = httpServer.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
			if s.isClose {
				return
			}
			log.Fatal("Server|Start|%s", err.Error())
		}
	})
}

// stop http server
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
