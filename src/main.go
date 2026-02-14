package main

import (
	"context"
	"log"

	"github.com/pasca-l/database-mcp/server"
)

func main() {
	s := server.NewDatabaseMCPServer()
	defer func() {
		if err := s.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	if err := s.Run(context.Background(), server.NewTransport()); err != nil {
		log.Printf("error running server: %v", err)
		return
	}
}
