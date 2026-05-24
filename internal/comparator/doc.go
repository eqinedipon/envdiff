// Package comparator provides functionality for comparing parsed .env
// files across multiple environments.
//
// It accepts a map of environment names to their key-value pairs and
// produces a structured Result that describes:
//
//   - Keys that are present in all environments with matching values
//   - Keys that are missing from one or more environments
//   - Keys whose values differ across environments (mismatches)
//
// Example usage:
//
//	envs := map[string]map[string]string{
//		"production": {"DB_HOST": "prod.db", "PORT": "5432"},
//		"staging":    {"DB_HOST": "stage.db", "PORT": "5432"},
//	}
//	result := comparator.Compare(envs)
//
The Result type is consumed by the reporter package to render output
in text, JSON, or Markdown formats.
package comparator
