// Package differ provides utilities for computing diffs between
// two sets of environment variable maps, producing structured change sets.
package differ

import "sort"

// ChangeKind describes the type of change detected between two env maps.
type ChangeKind string

const (
	Added    ChangeKind = "added"
	Removed  ChangeKind = "removed"
	Modified ChangeKind = "modified"
	Unchanged ChangeKind = "unchanged"
)

// Change represents a single key-level difference between two env maps.
type Change struct {
	Key      string
	Kind     ChangeKind
	OldValue string
	NewValue string
}

// Diff computes the difference between a base env map and a target env map.
// It returns a slice of Change entries sorted by key.
func Diff(base, target map[string]string) []Change {
	seen := make(map[string]bool)
	var changes []Change

	for k, baseVal := range base {
		seen[k] = true
		if targetVal, ok := target[k]; ok {
			if baseVal == targetVal {
				changes = append(changes, Change{Key: k, Kind: Unchanged, OldValue: baseVal, NewValue: targetVal})
			} else {
				changes = append(changes, Change{Key: k, Kind: Modified, OldValue: baseVal, NewValue: targetVal})
			}
		} else {
			changes = append(changes, Change{Key: k, Kind: Removed, OldValue: baseVal})
		}
	}

	for k, targetVal := range target {
		if !seen[k] {
			changes = append(changes, Change{Key: k, Kind: Added, NewValue: targetVal})
		}
	}

	sort.Slice(changes, func(i, j int) bool {
		return changes[i].Key < changes[j].Key
	})

	return changes
}

// Summary holds aggregate counts of each change kind.
type Summary struct {
	Added     int
	Removed   int
	Modified  int
	Unchanged int
}

// Summarize counts the kinds of changes in a slice of Change entries.
func Summarize(changes []Change) Summary {
	var s Summary
	for _, c := range changes {
		switch c.Kind {
		case Added:
			s.Added++
		case Removed:
			s.Removed++
		case Modified:
			s.Modified++
		case Unchanged:
			s.Unchanged++
		}
	}
	return s
}
