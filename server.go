package flexmem

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/eaciit/toolkit"
)

type Server struct {
	listener net.Listener
	log      *toolkit.LogEngine
}

func (s *Server) SetLogger(logger *toolkit.LogEngine) *Server {
	s.log = logger
	return s
}

func (s *Server) Start(host string) error {
	if s.log == nil {
		s.log = toolkit.NewLogEngine(true, false, "", "", "")
	}

	r := rpc.NewServer()
	p := new(RpcProxy).setLogger(s.log)
	RegisterRpcObjectToProxy(p, newKvdb())
	if err := r.Register(p); err != nil {
		return fmt.Errorf("start server fail. %s", err.Error())
	}

	l, err := net.Listen("tcp", host)
	if err != nil {
		return fmt.Errorf("start server fail. %s", err.Error())
	}
	s.listener = l

	go func() {
		r.Accept(l)
	}()

	s.log.Info("Server started")
	return nil
}

func (s *Server) Stop() {
	s.log.Info("Server is closed")
	go func() {
		s.listener.Close()
	}()
}
