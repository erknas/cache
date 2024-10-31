package main

import (
	"fmt"
	"log"
	"net"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	opts  ServerOpts
	cache Cacher
}

func NewServer(opts ServerOpts, cache Cacher) *Server {
	return &Server{
		opts:  opts,
		cache: cache,
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.opts.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}

	log.Printf("server running on port [%s]", s.opts.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("acception error: %s\n", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read error: %s\n", err)
			break
		}

		msg := buf[:n]
		fmt.Println(string(msg))
	}
}
