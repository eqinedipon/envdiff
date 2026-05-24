package reporter

import (
	"testing"
)

func TestDefaultOptions_Defaults(t *testing.T) {
	opts := DefaultOptions()

	if opts.Format != FormatText {
		t.Errorf("expected default format %q, got %q", FormatText, opts.Format)
	}
	if opts.Redact {
		t.Error("expected Redact to be false by default")
	}
	if !opts.Color {
		t.Error("expected Color to be true by default")
	}
	if opts.ShowMatches {
		t.Error("expected ShowMatches to be false by default")
	}
}

func TestFormat_Constants(t *testing.T) {
	cases := []struct {
		name   string
		format Format
	}{
		{"text", FormatText},
		{"json", FormatJSON},
		{"markdown", FormatMarkdown},
	}

	seen := map[Format]bool{}
	for _, tc := range cases {
		if tc.format == "" {
			t.Errorf("format constant %q must not be empty", tc.name)
		}
		if seen[tc.format] {
			t.Errorf("duplicate format value %q for %q", tc.format, tc.name)
		}
		seen[tc.format] = true
	}
}

func TestOptions_RedactToggle(t *testing.T) {
	opts := DefaultOptions()
	opts.Redact = true

	if !opts.Redact {
		t.Error("expected Redact to be true after setting")
	}
}

func TestOptions_ShowMatchesToggle(t *testing.T) {
	opts := DefaultOptions()
	opts.ShowMatches = true

	if !opts.ShowMatches {
		t.Error("expected ShowMatches to be true after setting")
	}
}

func TestOptions_ColorToggle(t *testing.T) {
	opts := DefaultOptions()
	opts.Color = false

	if opts.Color {
		t.Error("expected Color to be false after disabling")
	}
}
