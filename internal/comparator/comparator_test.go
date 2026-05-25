package comparator_test

import (
	"testing"

	"github.com/yourorg/envdiff/internal/comparator"
)

func TestCompare_AllMatch(t *testing.T) {
	envs := map[string]map[string]string{
		"dev":  {"HOST": "localhost", "PORT": "8080"},
		"prod": {"HOST": "localhost", "PORT": "8080"},
	}
	res := comparator.Compare(envs)
	for _, kr := range res.Keys {
		if kr.Status != comparator.StatusMatch {
			t.Errorf("key %q: expected match, got %s", kr.Key, kr.Status)
		}
	}
}

func TestCompare_MissingKey(t *testing.T) {
	envs := map[string]map[string]string{
		"dev":  {"HOST": "localhost", "SECRET": "abc"},
		"prod": {"HOST": "localhost"},
	}
	res := comparator.Compare(envs)
	found := false
	for _, kr := range res.Keys {
		if kr.Key == "SECRET" {
			found = true
			if kr.Status != comparator.StatusMissing {
				t.Errorf("SECRET: expected missing, got %s", kr.Status)
			}
			if kr.Values["prod"] != "" {
				t.Errorf("SECRET prod value should be empty string")
			}
		}
	}
	if !found {
		t.Error("SECRET key not found in results")
	}
}

func TestCompare_ValueMismatch(t *testing.T) {
	envs := map[string]map[string]string{
		"dev":  {"DB_URL": "postgres://localhost/dev"},
		"prod": {"DB_URL": "postgres://prod-host/prod"},
	}
	res := comparator.Compare(envs)
	for _, kr := range res.Keys {
		if kr.Key == "DB_URL" && kr.Status != comparator.StatusMismatch {
			t.Errorf("DB_URL: expected mismatch, got %s", kr.Status)
		}
	}
}

func TestCompare_EnvironmentsSorted(t *testing.T) {
	envs := map[string]map[string]string{
		"staging": {"A": "1"},
		"dev":     {"A": "1"},
		"prod":    {"A": "1"},
	}
	res := comparator.Compare(envs)
	expected := []string{"dev", "prod", "staging"}
	for i, e := range res.Environments {
		if e != expected[i] {
			t.Errorf("env[%d]: expected %s, got %s", i, expected[i], e)
		}
	}
}

func TestCompare_EmptyEnvs(t *testing.T) {
	envs := map[string]map[string]string{
		"dev":  {},
		"prod": {},
	}
	res := comparator.Compare(envs)
	if len(res.Keys) != 0 {
		t.Errorf("expected no keys, got %d", len(res.Keys))
	}
}

func TestCompare_KeysSorted(t *testing.T) {
	envs := map[string]map[string]string{
		"dev":  {"ZEBRA": "1", "ALPHA": "1", "MANGO": "1"},
		"prod": {"ZEBRA": "1", "ALPHA": "1", "MANGO": "1"},
	}
	res := comparator.Compare(envs)
	expected := []string{"ALPHA", "MANGO", "ZEBRA"}
	if len(res.Keys) != len(expected) {
		t.Fatalf("expected %d keys, got %d", len(expected), len(res.Keys))
	}
	for i, kr := range res.Keys {
		if kr.Key != expected[i] {
			t.Errorf("keys[%d]: expected %s, got %s", i, expected[i], kr.Key)
		}
	}
}
