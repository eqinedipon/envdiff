package reporter

import (
	"strings"
	"testing"

	"github.com/user/envdiff/internal/comparator"
)

func TestReport_MarkdownNoMismatch(t *testing.T) {
	result := makeResult(
		[]string{"staging", "production"},
		nil,
		nil,
	)
	var buf strings.Builder
	Report(&buf, result, Options{Format: "markdown"})
	out := buf.String()

	if !strings.Contains(out, "# envdiff Report") {
		t.Error("expected markdown heading")
	}
	if !strings.Contains(out, "All keys match") {
		t.Error("expected all-match message")
	}
}

func TestReport_MarkdownMissingKey(t *testing.T) {
	result := makeResult(
		[]string{"dev", "prod"},
		[]comparator.MissingKey{
			{Key: "SECRET", PresentIn: []string{"prod"}, AbsentFrom: []string{"dev"}},
		},
		nil,
	)
	var buf strings.Builder
	Report(&buf, result, Options{Format: "markdown"})
	out := buf.String()

	if !strings.Contains(out, "## Missing Keys") {
		t.Error("expected missing keys section")
	}
	if !strings.Contains(out, "`SECRET`") {
		t.Error("expected SECRET key in output")
	}
	if !strings.Contains(out, "prod") || !strings.Contains(out, "dev") {
		t.Error("expected environment names in output")
	}
}

func TestReport_MarkdownMismatch(t *testing.T) {
	result := makeResult(
		[]string{"dev", "prod"},
		nil,
		[]comparator.Mismatch{
			{Key: "DB_URL", Values: map[string]string{"dev": "localhost", "prod": "db.prod.example.com"}},
		},
	)
	var buf strings.Builder
	Report(&buf, result, Options{Format: "markdown"})
	out := buf.String()

	if !strings.Contains(out, "## Value Mismatches") {
		t.Error("expected mismatches section")
	}
	if !strings.Contains(out, "`DB_URL`") {
		t.Error("expected DB_URL key")
	}
	if !strings.Contains(out, "localhost") {
		t.Error("expected dev value")
	}
}

func TestReport_MarkdownRedact(t *testing.T) {
	result := makeResult(
		[]string{"dev", "prod"},
		nil,
		[]comparator.Mismatch{
			{Key: "API_KEY", Values: map[string]string{"dev": "abc123", "prod": "xyz789"}},
		},
	)
	var buf strings.Builder
	Report(&buf, result, Options{Format: "markdown", Redact: true})
	out := buf.String()

	if strings.Contains(out, "abc123") || strings.Contains(out, "xyz789") {
		t.Error("expected values to be redacted")
	}
	if !strings.Contains(out, "***") {
		t.Error("expected redaction placeholder")
	}
}
