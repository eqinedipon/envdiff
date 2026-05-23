// Command envdiff compares .env files across environments and highlights
// missing or mismatched keys.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/user/envdiff/internal/comparator"
	"github.com/user/envdiff/internal/parser"
	"github.com/user/envdiff/internal/reporter"
)

func main() {
	format := flag.String("format", "text", "Output format: text or json")
	noColor := flag.Bool("no-color", false, "Disable colored output")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: envdiff [flags] <file1> <file2> [file3...]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  envdiff .env.development .env.production\n")
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "error: at least two .env files are required")
		flag.Usage()
		os.Exit(1)
	}

	envs := make(map[string]map[string]string, len(args))
	for _, path := range args {
		name := envName(path)
		parsed, err := parser.ParseFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", path, err)
			os.Exit(1)
		}
		envs[name] = parsed
	}

	result := comparator.Compare(envs)

	opts := reporter.DefaultOptions()
	opts.Format = *format
	opts.NoColor = *noColor

	if err := reporter.Report(os.Stdout, result, opts); err != nil {
		fmt.Fprintf(os.Stderr, "error generating report: %v\n", err)
		os.Exit(1)
	}

	if result.HasDiff() {
		os.Exit(2)
	}
}

// envName derives a short environment label from a file path.
// e.g. ".env.production" -> "production", ".env" -> ".env"
func envName(path string) string {
	// Strip leading directory components
	base := path
	if idx := strings.LastIndex(path, "/"); idx >= 0 {
		base = path[idx+1:]
	}
	// Strip ".env." prefix if present
	if strings.HasPrefix(base, ".env.") {
		return strings.TrimPrefix(base, ".env.")
	}
	return base
}
