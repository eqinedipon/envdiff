package comparator

import "strings"

// Filter holds options for filtering comparison results.
type Filter struct {
	// IgnoreKeys is a list of exact key names to exclude from comparison.
	IgnoreKeys []string
	// IgnorePrefixes is a list of key prefixes to exclude from comparison.
	IgnorePrefixes []string
}

// NewFilter creates a Filter with the given ignore keys and prefixes.
func NewFilter(ignoreKeys, ignorePrefixes []string) Filter {
	return Filter{
		IgnoreKeys:     ignoreKeys,
		IgnorePrefixes: ignorePrefixes,
	}
}

// ShouldIgnore returns true if the given key matches any ignore rule.
func (f Filter) ShouldIgnore(key string) bool {
	for _, k := range f.IgnoreKeys {
		if k == key {
			return true
		}
	}
	for _, prefix := range f.IgnorePrefixes {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}
	return false
}

// ApplyFilter removes ignored keys from each environment map and returns new copies.
func ApplyFilter(envs map[string]map[string]string, f Filter) map[string]map[string]string {
	if len(f.IgnoreKeys) == 0 && len(f.IgnorePrefixes) == 0 {
		return envs
	}
	filtered := make(map[string]map[string]string, len(envs))
	for env, pairs := range envs {
		copy := make(map[string]string, len(pairs))
		for k, v := range pairs {
			if !f.ShouldIgnore(k) {
				copy[k] = v
			}
		}
		filtered[env] = copy
	}
	return filtered
}
