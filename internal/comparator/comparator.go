// Package comparator provides functionality for comparing parsed .env files
// across multiple environments, identifying missing keys and value mismatches.
package comparator

// KeyStatus represents the comparison status of a key across environments.
type KeyStatus string

const (
	// StatusMissing indicates the key is absent in one or more environments.
	StatusMissing KeyStatus = "missing"
	// StatusMismatch indicates the key exists in all environments but values differ.
	StatusMismatch KeyStatus = "mismatch"
	// StatusMatch indicates the key exists with the same value in all environments.
	StatusMatch KeyStatus = "match"
)

// KeyResult holds the comparison result for a single key.
type KeyResult struct {
	Key    string
	Status KeyStatus
	// Values maps environment name to the value found (empty string if missing).
	Values map[string]string
}

// Result holds the full comparison output for all keys across environments.
type Result struct {
	Environments []string
	Keys         []KeyResult
}

// Compare takes a map of environment name -> parsed key/value pairs and
// returns a Result describing matches, mismatches, and missing keys.
func Compare(envs map[string]map[string]string) Result {
	envNames := sortedKeys(envs)

	// Collect the union of all keys.
	allKeys := map[string]struct{}{}
	for _, pairs := range envs {
		for k := range pairs {
			allKeys[k] = struct{}{}
		}
	}

	var results []KeyResult
	for _, key := range sortedStringSlice(allKeys) {
		values := make(map[string]string, len(envNames))
		presentCount := 0
		firstVal := ""
		allSame := true

		for _, env := range envNames {
			v, ok := envs[env][key]
			if ok {
				presentCount++
				if firstVal == "" && presentCount == 1 {
					firstVal = v
				} else if v != firstVal {
					allSame = false
				}
				values[env] = v
			} else {
				values[env] = ""
				allSame = false
			}
		}

		var status KeyStatus
		switch {
		case presentCount < len(envNames):
			status = StatusMissing
		case !allSame:
			status = StatusMismatch
		default:
			status = StatusMatch
		}

		results = append(results, KeyResult{Key: key, Status: status, Values: values})
	}

	return Result{Environments: envNames, Keys: results}
}

// sortedKeys returns sorted keys of a map[string]map[string]string.
func sortedKeys(m map[string]map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return sortStringSlice(keys)
}

// sortedStringSlice converts a set to a sorted slice.
func sortedStringSlice(set map[string]struct{}) []string {
	slice := make([]string, 0, len(set))
	for k := range set {
		slice = append(slice, k)
	}
	return sortStringSlice(slice)
}

func sortStringSlice(s []string) []string {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
	return s
}
