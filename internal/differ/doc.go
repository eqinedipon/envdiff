// Package differ computes key-level differences between two environment
// variable maps, categorising each key as added, removed, modified, or
// unchanged.
//
// # Overview
//
// Use [Diff] to obtain a sorted slice of [Change] values describing every
// key present in either the base or the target map. Use [Summarize] to
// obtain aggregate counts of each [ChangeKind].
//
// # Example
//
//	base   := map[string]string{"HOST": "localhost", "PORT": "5432"}
//	target := map[string]string{"HOST": "db.prod",   "PORT": "5432", "SSL": "true"}
//
//	changes := differ.Diff(base, target)
//	summary := differ.Summarize(changes)
//	// summary → {Added:1, Removed:0, Modified:1, Unchanged:1}
package differ
