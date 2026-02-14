# Database MCP
MCP server for database connecting and querying.

## Screenshots

## Requirements
- Go 1.26

## Usage
1. Build the MCP server binary.
```bash
$ make build
```

2. Configure the binary path in the MCP server settings for any LLM service.
- for Claude Code, modify `.claude.json`
```json
{
  "mcpServers": {
    "database-mcp": {
      "command": "/ABSOLUTE_PATH_TO_BINARY/database-mcp"
    }
  }
}
```

- for GitHub Copilot in VSCode, modify `settings.json`
```json
{
  "mcp": {
    "servers": {
      "database-mcp": {
        "command": "/ABSOLUTE_PATH_TO_BINARY/database-mcp"
      }
    }
  }
}
```
