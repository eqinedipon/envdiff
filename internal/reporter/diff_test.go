package reporter

import (
	"strings"
	"testing"
)

func TestReportDiff_AddedKey(t *testing.T) {
	base := map[string]string{"A": "1"}
	target := map[string]string{"A": "1", "B": "2"}
	var buf strings.Builder
	opts := DefaultOptions()
	opts.Color = false
	if err := ReportDiff(&buf, base, target, ".env.base", ".env.target", opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "+ B=2") {
		t.Errorf("expected added key in output, got:\n%s", out)
	}
}

func TestReportDiff_RemovedKey(t *testing.T) {
	base := map[string]string{"A": "1", "B": "2"}
	target := map[string]string{"A": "1"}
	var buf strings.Builder
	opts := DefaultOptions()
	opts.Color = false
	_ = ReportDiff(&buf, base, target, "base", "target", opts)
	out := buf.String()
	if !strings.Contains(out, "- B=2") {
		t.Errorf("expected removed key in output, got:\n%s", out)
	}
}

func TestReportDiff_ModifiedKey(t *testing.T) {
	base := map[string]string{"HOST": "localhost"}
	target := map[string]string{"HOST": "prod.db"}
	var buf strings.Builder
	opts := DefaultOptions()
	opts.Color = false
	_ = ReportDiff(&buf, base, target, "base", "target", opts)
	out := buf.String()
	if !strings.Contains(out, "~ HOST") {
		t.Errorf("expected modified key in output, got:\n%s", out)
	}
}

func TestReportDiff_Redact(t *testing.T) {
	base := map[string]string{"SECRET": "hunter2"}
	target := map[string]string{"SECRET": "newpass"}
	var buf strings.Builder
	opts := DefaultOptions()
	opts.Color = false
	opts.Redact = true
	_ = ReportDiff(&buf, base, target, "base", "target", opts)
	out := buf.String()
	if strings.Contains(out, "hunter2") || strings.Contains(out, "newpass") {
		t.Errorf("expected values to be redacted, got:\n%s", out)
	}
	if !strings.Contains(out, "***") {
		t.Errorf("expected redacted placeholder, got:\n%s", out)
	}
}

func TestReportDiff_SummaryLine(t *testing.T) {
	base := map[string]string{"A": "1", "B": "old"}
	target := map[string]string{"B": "new", "C": "3"}
	var buf strings.Builder
	opts := DefaultOptions()
	opts.Color = false
	_ = ReportDiff(&buf, base, target, "base", "target", opts)
	out := buf.String()
	if !strings.Contains(out, "Summary:") {
		t.Errorf("expected summary line, got:\n%s", out)
	}
}
