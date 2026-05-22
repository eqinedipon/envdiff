// Package parser provides utilities for reading and parsing .env files
// into structured key-value maps.
//
// Supported syntax:
//   - KEY=VALUE
//   - KEY="VALUE"  (double-quoted values)
//   - KEY='VALUE'  (single-quoted values)
//   - Lines starting with '#' are treated as comments and ignored.
//   - Blank lines are ignored.
//
// Example usage:
//
//	env, err := parser.ParseFile(".env.production")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(env["APP_ENV"])
package parser
