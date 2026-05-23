// Package reporter formats and outputs the results of an env file comparison.
package reporter

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/user/envdiff/internal/comparator"
)

// Format represents the output format for the report.
type Format string

const (
	FormatText Format = "text"
	FormatJSON  Format = "json"
)

// Options configures the reporter output.
type Options struct {
	Format  Format
	Writer  io.Writer
	NoColor bool
}

// DefaultOptions returns sensible defaults for the reporter.
func DefaultOptions() Options {
	return Options{
		Format:  FormatText,
		Writer:  os.Stdout,
		NoColor: false,
	}
}

// Report writes a human-readable diff report to the configured writer.
func Report(result comparator.Result, opts Options) error {
	if opts.Writer == nil {
		opts.Writer = os.Stdout
	}

	switch opts.Format {
	case FormatJSON:
		return reportJSON(result, opts.Writer)
	default:
		return reportText(result, opts.Writer, opts.NoColor)
	}
}

func reportText(result comparator.Result, w io.Writer, noColor bool) error {
	if len(result.Diffs) == 0 {
		fmt.Fprintln(w, "✓ All environments match.")
		return nil
	}

	fmt.Fprintf(w, "Found %d difference(s):\n\n", len(result.Diffs))

	for _, diff := range result.Diffs {
		switch diff.Type {
		case comparator.DiffMissing:
			line := fmt.Sprintf("  [MISSING] key %q absent in: %s", diff.Key, strings.Join(diff.Environments, ", "))
			fmt.Fprintln(w, colorize(line, colorYellow, noColor))
		case comparator.DiffMismatch:
			fmt.Fprintf(w, "  [MISMATCH] key %q:\n", diff.Key)
			for env, val := range diff.Values {
				line := fmt.Sprintf("    %s = %q", env, val)
				fmt.Fprintln(w, colorize(line, colorRed, noColor))
			}
		}
	}
	return nil
}

const (
	colorReset  = "\033[0m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

func colorize(s, color string, noColor bool) string {
	if noColor {
		return s
	}
	return color + s + colorReset
}
