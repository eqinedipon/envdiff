package reporter

import (
	"fmt"
	"io"
	"strings"

	"github.com/yourorg/envdiff/internal/differ"
)

// ReportDiff writes a human-readable diff between base and target env maps
// to w, respecting the supplied Options (color, redact).
func ReportDiff(w io.Writer, base, target map[string]string, baseLabel, targetLabel string, opts Options) error {
	changes := differ.Diff(base, target)
	summary := differ.Summarize(changes)

	fmt.Fprintf(w, "Diff: %s → %s\n", baseLabel, targetLabel)
	fmt.Fprintln(w, strings.Repeat("─", 40))

	for _, c := range changes {
		switch c.Kind {
		case differ.Added:
			val := redactIfNeeded(c.NewValue, opts)
			fmt.Fprintf(w, "%s %s=%s\n", colorize("+", "green", opts.Color), c.Key, val)
		case differ.Removed:
			val := redactIfNeeded(c.OldValue, opts)
			fmt.Fprintf(w, "%s %s=%s\n", colorize("-", "red", opts.Color), c.Key, val)
		case differ.Modified:
			old := redactIfNeeded(c.OldValue, opts)
			new := redactIfNeeded(c.NewValue, opts)
			fmt.Fprintf(w, "%s %s: %s → %s\n", colorize("~", "yellow", opts.Color), c.Key, old, new)
		case differ.Unchanged:
			if opts.ShowMatches {
				val := redactIfNeeded(c.OldValue, opts)
				fmt.Fprintf(w, "  %s=%s\n", c.Key, val)
			}
		}
	}

	fmt.Fprintln(w, strings.Repeat("─", 40))
	fmt.Fprintf(w, "Summary: +%d added  -%d removed  ~%d modified  =%d unchanged\n",
		summary.Added, summary.Removed, summary.Modified, summary.Unchanged)
	return nil
}

func redactIfNeeded(val string, opts Options) string {
	if opts.Redact && val != "" {
		return "***"
	}
	return val
}
