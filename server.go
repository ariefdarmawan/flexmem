package flexmem

import (
	"fmt"

	"github.com/ariefdarmawan/rpchub/hubserver"

	"github.com/eaciit/toolkit"
)

type Server struct {
	//listener net.Listener
	log *toolkit.LogEngine
	hs  *hubserver.Server
}

func (s *Server) SetLogger(logger *toolkit.LogEngine) *Server {
	s.log = logger
	return s
}

func (s *Server) Start(host string) error {
	if s.log == nil {
		s.log = toolkit.NewLogEngine(true, false, "", "", "")
	}

	hs := hubserver.NewServer().SetLog(s.log)
	hs.Register(NewKvDB())
	if err := hs.Start(host); err != nil {
		return fmt.Errorf("unable to start flexmem server. %s", err.Error())
	}
	s.hs = hs
	s.log.Info("Server started")
	return nil
}

func (s *Server) Stop() {
	s.log.Info("Server is closed")

	go func() {
		if s.hs != nil {
			s.hs.Stop()
		}
	}()
}
