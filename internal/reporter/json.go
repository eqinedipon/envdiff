package reporter

import (
	"encoding/json"
	"io"

	"github.com/user/envdiff/internal/comparator"
)

// jsonDiff is the JSON-serialisable representation of a single diff entry.
type jsonDiff struct {
	Type         string            `json:"type"`
	Key          string            `json:"key"`
	Environments []string          `json:"environments,omitempty"`
	Values       map[string]string `json:"values,omitempty"`
}

// jsonReport is the top-level JSON output structure.
type jsonReport struct {
	Match bool       `json:"match"`
	Diffs []jsonDiff `json:"diffs"`
}

func reportJSON(result comparator.Result, w io.Writer) error {
	report := jsonReport{
		Match: len(result.Diffs) == 0,
		Diffs: make([]jsonDiff, 0, len(result.Diffs)),
	}

	for _, d := range result.Diffs {
		jd := jsonDiff{
			Key: d.Key,
		}
		switch d.Type {
		case comparator.DiffMissing:
			jd.Type = "missing"
			jd.Environments = d.Environments
		case comparator.DiffMismatch:
			jd.Type = "mismatch"
			jd.Values = d.Values
		}
		report.Diffs = append(report.Diffs, jd)
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(report)
}
