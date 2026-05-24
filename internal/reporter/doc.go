// Package reporter provides formatting and output utilities for envdiff results.
//
// It supports multiple output formats including plain text, JSON, and Markdown.
// Each format can be configured via Options, which control behaviour such as
// redacting sensitive values, enabling colour output, and toggling match display.
//
// Usage:
//
//	result := comparator.Compare(envs)
//	opts := reporter.DefaultOptions()
//	opts.Format = reporter.FormatJSON
//	reporter.Report(os.Stdout, result, opts)
//
// Supported formats:
//
//	- FormatText     plain human-readable text (default)
//	- FormatJSON     machine-readable JSON
//	- FormatMarkdown GitHub-flavoured Markdown table
package reporter
