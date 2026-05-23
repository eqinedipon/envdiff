package reporter

import (
	"fmt"
	"io"
	"strings"

	"github.com/user/envdiff/internal/comparator"
)

// reportMarkdown writes a Markdown-formatted diff report to w.
func reportMarkdown(w io.Writer, result comparator.Result, opts Options) {
	fmt.Fprintf(w, "# envdiff Report\n\n")

	envs := result.Environments
	if len(envs) == 0 {
		fmt.Fprintf(w, "_No environments compared._\n")
		return
	}

	// Summary line
	fmt.Fprintf(w, "**Environments:** %s\n\n", strings.Join(envs, ", "))

	if len(result.Missing) == 0 && len(result.Mismatches) == 0 {
		fmt.Fprintf(w, "✅ All keys match across environments.\n")
		return
	}

	if len(result.Missing) > 0 {
		fmt.Fprintf(w, "## Missing Keys\n\n")
		fmt.Fprintf(w, "| Key | Present In | Absent From |\n")
		fmt.Fprintf(w, "|-----|-----------|------------|\n")
		for _, m := range result.Missing {
			present := strings.Join(m.PresentIn, ", ")
			absent := strings.Join(m.AbsentFrom, ", ")
			fmt.Fprintf(w, "| `%s` | %s | %s |\n", m.Key, present, absent)
		}
		fmt.Fprintf(w, "\n")
	}

	if len(result.Mismatches) > 0 {
		fmt.Fprintf(w, "## Value Mismatches\n\n")
		for _, mm := range result.Mismatches {
			fmt.Fprintf(w, "### `%s`\n\n", mm.Key)
			fmt.Fprintf(w, "| Environment | Value |\n")
			fmt.Fprintf(w, "|-------------|-------|\n")
			for _, env := range envs {
				val, ok := mm.Values[env]
				if !ok {
					val = "_(missing)_"
				} else if opts.Redact {
					val = "***"
				}
				fmt.Fprintf(w, "| %s | `%s` |\n", env, val)
			}
			fmt.Fprintf(w, "\n")
		}
	}
}
