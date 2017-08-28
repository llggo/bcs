package conf

import (
	"context"
	"fmt"
	"time"
)

type ServerConfig struct {
	Name         string
	Host         string
	Port         int
	ShutdownWait int
}

func (s *ServerConfig) Addr() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

func (s *ServerConfig) AddrName() string {
	return s.Name
}

func (s ServerConfig) String() string {
	return fmt.Sprintf("server:host=%v,port=%v", s.Host, s.Port)
}

func (s *ServerConfig) Wait() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(s.ShutdownWait)*time.Second)
}
