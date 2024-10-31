package main

import (
	"log"
	"net"
)

func main() {
	var (
		opts = ServerOpts{
			ListenAddr: ":3000",
			IsLeader:   true,
		}
		cache = NewCache()
	)

	go func() {
		conn, err := net.Dial("tcp", opts.ListenAddr)
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("test 123"))
	}()

	srv := NewServer(opts, cache)
	srv.Run()
}
