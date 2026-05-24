package comparator

import (
	"testing"
)

func TestShouldIgnore_ExactKey(t *testing.T) {
	f := NewFilter([]string{"SECRET", "TOKEN"}, nil)
	if !f.ShouldIgnore("SECRET") {
		t.Error("expected SECRET to be ignored")
	}
	if !f.ShouldIgnore("TOKEN") {
		t.Error("expected TOKEN to be ignored")
	}
	if f.ShouldIgnore("DATABASE_URL") {
		t.Error("expected DATABASE_URL not to be ignored")
	}
}

func TestShouldIgnore_Prefix(t *testing.T) {
	f := NewFilter(nil, []string{"AWS_", "INTERNAL_"})
	if !f.ShouldIgnore("AWS_ACCESS_KEY") {
		t.Error("expected AWS_ACCESS_KEY to be ignored")
	}
	if !f.ShouldIgnore("INTERNAL_FLAG") {
		t.Error("expected INTERNAL_FLAG to be ignored")
	}
	if f.ShouldIgnore("APP_ENV") {
		t.Error("expected APP_ENV not to be ignored")
	}
}

func TestShouldIgnore_EmptyFilter(t *testing.T) {
	f := NewFilter(nil, nil)
	if f.ShouldIgnore("ANYTHING") {
		t.Error("empty filter should not ignore any key")
	}
}

func TestApplyFilter_RemovesKeys(t *testing.T) {
	envs := map[string]map[string]string{
		"dev": {"APP_ENV": "development", "SECRET": "abc", "AWS_KEY": "key1"},
		"prod": {"APP_ENV": "production", "SECRET": "xyz", "AWS_KEY": "key2"},
	}
	f := NewFilter([]string{"SECRET"}, []string{"AWS_"})
	result := ApplyFilter(envs, f)

	for env, pairs := range result {
		if _, ok := pairs["SECRET"]; ok {
			t.Errorf("env %s: SECRET should have been filtered out", env)
		}
		if _, ok := pairs["AWS_KEY"]; ok {
			t.Errorf("env %s: AWS_KEY should have been filtered out", env)
		}
		if _, ok := pairs["APP_ENV"]; !ok {
			t.Errorf("env %s: APP_ENV should remain", env)
		}
	}
}

func TestApplyFilter_NoOpWhenEmpty(t *testing.T) {
	envs := map[string]map[string]string{
		"dev": {"KEY": "val"},
	}
	f := NewFilter(nil, nil)
	result := ApplyFilter(envs, f)
	if result["dev"]["KEY"] != "val" {
		t.Error("expected KEY to remain unchanged")
	}
}

func TestApplyFilter_OriginalUnmodified(t *testing.T) {
	original := map[string]map[string]string{
		"dev": {"SECRET": "s", "APP": "a"},
	}
	f := NewFilter([]string{"SECRET"}, nil)
	ApplyFilter(original, f)
	if _, ok := original["dev"]["SECRET"]; !ok {
		t.Error("ApplyFilter should not modify the original map")
	}
}
