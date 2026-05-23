package reporter_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/user/envdiff/internal/comparator"
	"github.com/user/envdiff/internal/reporter"
)

func makeResult(diffs []comparator.Diff) comparator.Result {
	return comparator.Result{Diffs: diffs}
}

func TestReport_TextNoMismatch(t *testing.T) {
	var buf bytes.Buffer
	opts := reporter.Options{Format: reporter.FormatText, Writer: &buf, NoColor: true}
	if err := reporter.Report(makeResult(nil), opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "All environments match") {
		t.Errorf("expected match message, got: %s", buf.String())
	}
}

func TestReport_TextMissingKey(t *testing.T) {
	var buf bytes.Buffer
	opts := reporter.Options{Format: reporter.FormatText, Writer: &buf, NoColor: true}
	diffs := []comparator.Diff{
		{Type: comparator.DiffMissing, Key: "DB_HOST", Environments: []string{"staging"}},
	}
	if err := reporter.Report(makeResult(diffs), opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "MISSING") || !strings.Contains(out, "DB_HOST") {
		t.Errorf("expected MISSING/DB_HOST in output, got: %s", out)
	}
}

func TestReport_TextMismatch(t *testing.T) {
	var buf bytes.Buffer
	opts := reporter.Options{Format: reporter.FormatText, Writer: &buf, NoColor: true}
	diffs := []comparator.Diff{
		{Type: comparator.DiffMismatch, Key: "LOG_LEVEL", Values: map[string]string{"dev": "debug", "prod": "error"}},
	}
	if err := reporter.Report(makeResult(diffs), opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "MISMATCH") || !strings.Contains(out, "LOG_LEVEL") {
		t.Errorf("expected MISMATCH/LOG_LEVEL in output, got: %s", out)
	}
}

func TestReport_JSONFormat(t *testing.T) {
	var buf bytes.Buffer
	opts := reporter.Options{Format: reporter.FormatJSON, Writer: &buf}
	diffs := []comparator.Diff{
		{Type: comparator.DiffMissing, Key: "SECRET", Environments: []string{"prod"}},
	}
	if err := reporter.Report(makeResult(diffs), opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if out["match"] != false {
		t.Errorf("expected match=false")
	}
}

func TestReport_JSONMatch(t *testing.T) {
	var buf bytes.Buffer
	opts := reporter.Options{Format: reporter.FormatJSON, Writer: &buf}
	if err := reporter.Report(makeResult(nil), opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if out["match"] != true {
		t.Errorf("expected match=true")
	}
}
