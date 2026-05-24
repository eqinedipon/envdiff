package comparator

// KeyStatus represents the comparison status of a single key across environments.
type KeyStatus int

const (
	// StatusMatch indicates the key exists in all environments with the same value.
	StatusMatch KeyStatus = iota
	// StatusMissing indicates the key is absent from one or more environments.
	StatusMissing
	// StatusMismatch indicates the key exists in all environments but values differ.
	StatusMismatch
)

// String returns a human-readable label for a KeyStatus.
func (s KeyStatus) String() string {
	switch s {
	case StatusMatch:
		return "match"
	case StatusMissing:
		return "missing"
	case StatusMismatch:
		return "mismatch"
	default:
		return "unknown"
	}
}

// KeyResult holds the comparison outcome for a single key.
type KeyResult struct {
	// Key is the environment variable name.
	Key string
	// Status is the overall comparison status for this key.
	Status KeyStatus
	// Values maps each environment name to the value found (empty string if absent).
	Values map[string]string
	// PresentIn lists the environments where this key exists.
	PresentIn []string
	// MissingFrom lists the environments where this key is absent.
	MissingFrom []string
}

// Result is the top-level output of a comparison operation.
type Result struct {
	// Environments is the sorted list of environment names that were compared.
	Environments []string
	// Keys holds one KeyResult per unique key found across all environments.
	Keys []KeyResult
	// TotalKeys is the total number of unique keys discovered.
	TotalKeys int
	// MatchCount is the number of keys that matched across all environments.
	MatchCount int
	// MissingCount is the number of keys missing from at least one environment.
	MissingCount int
	// MismatchCount is the number of keys present everywhere but with differing values.
	MismatchCount int
}

// HasIssues returns true when the result contains any missing or mismatched keys.
func (r Result) HasIssues() bool {
	return r.MissingCount > 0 || r.MismatchCount > 0
}
