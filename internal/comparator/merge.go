package comparator

// MergeEnvs merges multiple environment maps into a single map.
// When a key exists in multiple environments, the value from the last
// environment in the list that contains the key is used.
// The merge also returns a report of which keys were overridden and by which env.
type MergeConflict struct {
	Key      string
	Original string
	Override string
	FromEnv  string
}

// MergeResult holds the merged key-value map and any conflicts encountered.
type MergeResult struct {
	Merged    map[string]string
	Conflicts []MergeConflict
}

// MergeEnvs merges the provided named environment maps into a single flat map.
// envs is a slice of (name, map) pairs represented as EnvEntry.
type EnvEntry struct {
	Name   string
	Values map[string]string
}

// MergeEnvs combines multiple env maps into one, recording conflicts where
// a key appears in more than one environment with a different value.
func MergeEnvs(envs []EnvEntry) MergeResult {
	merged := make(map[string]string)
	conflicts := []MergeConflict{}
	seen := make(map[string]string) // key -> env name that set it

	for _, entry := range envs {
		for k, v := range entry.Values {
			if existing, ok := merged[k]; ok {
				if existing != v {
					conflicts = append(conflicts, MergeConflict{
						Key:      k,
						Original: existing,
						Override: v,
						FromEnv:  entry.Name,
					})
					_ = seen[k] // already tracked
				}
			}
			merged[k] = v
			seen[k] = entry.Name
		}
	}

	return MergeResult{
		Merged:    merged,
		Conflicts: conflicts,
	}
}
