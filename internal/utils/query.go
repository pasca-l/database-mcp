package utils

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrNotSelectQuery    = errors.New("query must be a SELECT statement")
	ErrDangerousQuery    = errors.New("query contains potentially dangerous operations")
	ErrCommentNotAllowed = errors.New("comments are not allowed in queries")
)

var (
	// Pattern to check if query starts with SELECT or WITH (for CTEs).
	selectOrWithPattern = regexp.MustCompile(`(?i)^(SELECT|WITH)\s`)
	// Pattern to detect any SQL comments (both -- and /* */ style).
	commentPattern = regexp.MustCompile(`(?i)(--.*)|(/\*.*?\*/)`)
	// Pattern to detect dangerous SQL patterns.
	dangerousPatterns = regexp.MustCompile(`(?i)(;\s*(DROP|DELETE|UPDATE|INSERT|ALTER|CREATE|TRUNCATE|EXEC|EXECUTE))|(/\*.*\*/.*?(DROP|DELETE|UPDATE|INSERT|ALTER|CREATE|TRUNCATE))`)
)

func ValidateReadQuery(query string) error {
	trimmed := strings.TrimSpace(query)

	if !selectOrWithPattern.MatchString(trimmed) {
		return ErrNotSelectQuery
	}
	if commentPattern.MatchString(trimmed) {
		return ErrCommentNotAllowed
	}
	if dangerousPatterns.MatchString(trimmed) {
		return ErrDangerousQuery
	}
	return nil
}
