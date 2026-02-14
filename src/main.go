package main

import (
	"context"
	"log"

	"github.com/pasca-l/database-mcp/server"
)

func main() {
	s := server.NewDatabaseMCPServer()
	t := server.NewTransport()

	err := s.Run(context.Background(), t)
	if err != nil {
		log.Fatal(err)
	}
}
