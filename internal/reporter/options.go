// Package reporter provides output formatting for envdiff comparison results.
// This file defines the Options type and related constants used to configure
// how comparison results are rendered.
package reporter

// Format represents the output format for the report.
type Format string

const (
	// FormatText renders results as human-readable plain text with optional color.
	FormatText Format = "text"

	// FormatJSON renders results as a JSON document.
	FormatJSON Format = "json"

	// FormatMarkdown renders results as a Markdown table.
	FormatMarkdown Format = "markdown"
)

// Options controls the behaviour of the Report function.
type Options struct {
	// Format selects the output format. Defaults to FormatText.
	Format Format

	// Color enables ANSI color codes in text output. Has no effect for JSON
	// or Markdown formats.
	Color bool

	// Redact replaces actual values with "***" in the output. Useful when
	// sharing reports that may contain sensitive data.
	Redact bool

	// ShowMatching includes keys whose values match across all environments
	// in the output. By default only differences are shown.
	ShowMatching bool
}

// DefaultOptions returns an Options value with sensible defaults:
// plain text output, color enabled, values not redacted, and matching
// keys hidden.
func DefaultOptions() Options {
	return Options{
		Format:       FormatText,
		Color:        true,
		Redact:       false,
		ShowMatching: false,
	}
}

// Validate checks that the Format field contains a recognised value and
// returns an error if it does not.
func (o Options) Validate() error {
	switch o.Format {
	case FormatText, FormatJSON, FormatMarkdown:
		return nil
	default:
		return &UnknownFormatError{Format: string(o.Format)}
	}
}

// UnknownFormatError is returned by Options.Validate when an unrecognised
// output format is specified.
type UnknownFormatError struct {
	Format string
}

func (e *UnknownFormatError) Error() string {
	return "envdiff: unknown output format \"" + e.Format + "\": must be one of text, json, markdown"
}
