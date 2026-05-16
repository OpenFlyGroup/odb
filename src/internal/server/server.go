package server

type Config struct {
	DataDir    string
	ListenAddr string
}

type Server struct {
	config Config
}

func New(config Config) (*Server, error) {
	return &Server{config: config}, nil
}

func (s *Server) Run() error {
	select {}
}

func (s *Server) Shutdown() error {
	return nil
}
