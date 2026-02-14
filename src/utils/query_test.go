package utils

import (
	"testing"
)

func TestValidateReadQuery(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		wantErr error
	}{
		// Valid queries - various SELECT structures
		{
			name:    "simple SELECT",
			query:   "SELECT * FROM users",
			wantErr: nil,
		},
		{
			name:    "complex SELECT",
			query:   "SELECT u.name, o.total FROM users u JOIN orders o ON u.id = o.user_id WHERE u.id = 1",
			wantErr: nil,
		},
		{
			name:    "SELECT with subquery",
			query:   "SELECT * FROM (SELECT id FROM users) AS sub",
			wantErr: nil,
		},

		// Valid queries - WITH (CTE)
		{
			name:    "WITH CTE single",
			query:   "WITH temp AS (SELECT id FROM users) SELECT * FROM temp",
			wantErr: nil,
		},
		{
			name:    "WITH CTE multiple",
			query:   "WITH cte1 AS (SELECT id FROM users), cte2 AS (SELECT name FROM users) SELECT * FROM cte1 JOIN cte2",
			wantErr: nil,
		},

		// Valid queries - case insensitivity and whitespace handling
		{
			name:    "lowercase select",
			query:   "select * from users",
			wantErr: nil,
		},
		{
			name:    "lowercase with",
			query:   "with temp as (select id from users) select * from temp",
			wantErr: nil,
		},
		{
			name:    "leading whitespace",
			query:   "  \n\t  SELECT * FROM users",
			wantErr: nil,
		},
		{
			name:    "keywords in identifiers",
			query:   "SELECT user_select, dropdown, withholding FROM users",
			wantErr: nil,
		},

		// Invalid - non-SELECT statements
		{
			name:    "INSERT",
			query:   "INSERT INTO users (name) VALUES ('John')",
			wantErr: ErrNotSelectQuery,
		},
		{
			name:    "empty string",
			query:   "",
			wantErr: ErrNotSelectQuery,
		},
		{
			name:    "whitespace only",
			query:   "   \n\t  ",
			wantErr: ErrNotSelectQuery,
		},

		// Invalid - comments (both styles)
		{
			name:    "comment before SELECT",
			query:   "-- comment\nSELECT * FROM users",
			wantErr: ErrNotSelectQuery,
		},
		{
			name:    "inline comment --",
			query:   "SELECT * FROM users -- comment",
			wantErr: ErrCommentNotAllowed,
		},
		{
			name:    "inline comment /* */",
			query:   "SELECT * FROM users /* comment */",
			wantErr: ErrCommentNotAllowed,
		},

		// Invalid - SQL injection attempts (semicolon + dangerous operations)
		{
			name:    "injection DROP",
			query:   "SELECT * FROM users; DROP TABLE users",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection DELETE",
			query:   "SELECT * FROM users; DELETE FROM users",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection UPDATE",
			query:   "SELECT * FROM users; UPDATE users SET x=1",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection INSERT",
			query:   "SELECT * FROM users; INSERT INTO users VALUES (1)",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection ALTER",
			query:   "SELECT * FROM users; ALTER TABLE users ADD COLUMN x INT",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection CREATE",
			query:   "SELECT * FROM users; CREATE TABLE test (id INT)",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection TRUNCATE",
			query:   "SELECT * FROM users; TRUNCATE TABLE users",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection EXEC",
			query:   "SELECT * FROM users; EXEC sp_executesql",
			wantErr: ErrDangerousQuery,
		},

		// Invalid - injection with variations
		{
			name:    "injection with whitespace",
			query:   "SELECT * FROM users;   DROP TABLE users",
			wantErr: ErrDangerousQuery,
		},
		{
			name:    "injection case insensitive",
			query:   "SELECT * FROM users; drop table users",
			wantErr: ErrDangerousQuery,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateReadQuery(tt.query)
			if err != tt.wantErr {
				t.Errorf("ValidateReadQuery(%q) = %v, want %v", tt.query, err, tt.wantErr)
			}
		})
	}
}
