package differ_test

import (
	"testing"

	"github.com/yourorg/envdiff/internal/differ"
)

func TestDiff_Added(t *testing.T) {
	base := map[string]string{"A": "1"}
	target := map[string]string{"A": "1", "B": "2"}
	changes := differ.Diff(base, target)
	if len(changes) != 2 {
		t.Fatalf("expected 2 changes, got %d", len(changes))
	}
	if changes[0].Kind != differ.Unchanged || changes[0].Key != "A" {
		t.Errorf("expected A unchanged, got %+v", changes[0])
	}
	if changes[1].Kind != differ.Added || changes[1].Key != "B" {
		t.Errorf("expected B added, got %+v", changes[1])
	}
}

func TestDiff_Removed(t *testing.T) {
	base := map[string]string{"A": "1", "B": "2"}
	target := map[string]string{"A": "1"}
	changes := differ.Diff(base, target)
	if len(changes) != 2 {
		t.Fatalf("expected 2 changes, got %d", len(changes))
	}
	var found bool
	for _, c := range changes {
		if c.Key == "B" && c.Kind == differ.Removed {
			found = true
		}
	}
	if !found {
		t.Error("expected B to be removed")
	}
}

func TestDiff_Modified(t *testing.T) {
	base := map[string]string{"KEY": "old"}
	target := map[string]string{"KEY": "new"}
	changes := differ.Diff(base, target)
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].Kind != differ.Modified {
		t.Errorf("expected Modified, got %s", changes[0].Kind)
	}
	if changes[0].OldValue != "old" || changes[0].NewValue != "new" {
		t.Errorf("unexpected values: %+v", changes[0])
	}
}

func TestDiff_Empty(t *testing.T) {
	changes := differ.Diff(map[string]string{}, map[string]string{})
	if len(changes) != 0 {
		t.Errorf("expected no changes, got %d", len(changes))
	}
}

func TestDiff_SortedByKey(t *testing.T) {
	base := map[string]string{"Z": "1", "A": "2", "M": "3"}
	target := map[string]string{"Z": "1", "A": "2", "M": "3"}
	changes := differ.Diff(base, target)
	for i := 1; i < len(changes); i++ {
		if changes[i].Key < changes[i-1].Key {
			t.Errorf("changes not sorted at index %d: %s < %s", i, changes[i].Key, changes[i-1].Key)
		}
	}
}

func TestSummarize(t *testing.T) {
	base := map[string]string{"A": "1", "B": "old", "C": "3"}
	target := map[string]string{"B": "new", "C": "3", "D": "4"}
	changes := differ.Diff(base, target)
	s := differ.Summarize(changes)
	if s.Added != 1 {
		t.Errorf("expected 1 added, got %d", s.Added)
	}
	if s.Removed != 1 {
		t.Errorf("expected 1 removed, got %d", s.Removed)
	}
	if s.Modified != 1 {
		t.Errorf("expected 1 modified, got %d", s.Modified)
	}
	if s.Unchanged != 1 {
		t.Errorf("expected 1 unchanged, got %d", s.Unchanged)
	}
}
