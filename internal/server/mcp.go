package server

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pasca-l/database-mcp/internal/db"
	"github.com/pasca-l/database-mcp/internal/tools"
)

type DatabaseMCPServer struct {
	*mcp.Server
	conn *db.Connection
}

func NewDatabaseMCPServer() *DatabaseMCPServer {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "database-mcp",
			Version: "v0.1.0",
		},
		nil,
	)

	conn := db.NewConnection()

	connectTool := tools.NewConnectTool(conn)
	readQueryTool := tools.NewReadQueryTool(conn)
	writeQueryTool := tools.NewWriteQueryTool(conn)

	mcp.AddTool(server, connectTool.Tool, connectTool.Handler)
	mcp.AddTool(server, readQueryTool.Tool, readQueryTool.Handler)
	mcp.AddTool(server, writeQueryTool.Tool, writeQueryTool.Handler)

	return &DatabaseMCPServer{
		Server: server,
		conn:   conn,
	}
}

func (s *DatabaseMCPServer) Close() error {
	return s.conn.Close()
}
