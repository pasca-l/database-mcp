# Database MCP
MCP server for database connecting and querying.
> [!WARNING]
> Users assume full responsibility for all database operations performed with this tool. Developers of this project are not liable for any data loss, corruption, or unintended modifications.

## Screenshots
![](https://github.com/user-attachments/assets/639d329b-4ce5-4650-89f1-9ce4d6345670)

## Requirements
- mise v2026.2.11

## Usage
1. Build the MCP server binary.
```bash
$ mise build
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
