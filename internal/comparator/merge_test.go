package comparator

import (
	"testing"
)

func TestMergeEnvs_NoConflicts(t *testing.T) {
	envs := []EnvEntry{
		{Name: "base", Values: map[string]string{"A": "1", "B": "2"}},
		{Name: "prod", Values: map[string]string{"C": "3"}},
	}
	result := MergeEnvs(envs)
	if len(result.Conflicts) != 0 {
		t.Errorf("expected no conflicts, got %d", len(result.Conflicts))
	}
	if result.Merged["A"] != "1" || result.Merged["B"] != "2" || result.Merged["C"] != "3" {
		t.Errorf("unexpected merged values: %v", result.Merged)
	}
}

func TestMergeEnvs_WithConflict(t *testing.T) {
	envs := []EnvEntry{
		{Name: "base", Values: map[string]string{"HOST": "localhost"}},
		{Name: "prod", Values: map[string]string{"HOST": "prod.example.com"}},
	}
	result := MergeEnvs(envs)
	if len(result.Conflicts) != 1 {
		t.Fatalf("expected 1 conflict, got %d", len(result.Conflicts))
	}
	c := result.Conflicts[0]
	if c.Key != "HOST" {
		t.Errorf("expected conflict key HOST, got %s", c.Key)
	}
	if c.Original != "localhost" {
		t.Errorf("expected original localhost, got %s", c.Original)
	}
	if c.Override != "prod.example.com" {
		t.Errorf("expected override prod.example.com, got %s", c.Override)
	}
	if c.FromEnv != "prod" {
		t.Errorf("expected FromEnv prod, got %s", c.FromEnv)
	}
	if result.Merged["HOST"] != "prod.example.com" {
		t.Errorf("expected merged HOST to be prod.example.com, got %s", result.Merged["HOST"])
	}
}

func TestMergeEnvs_EmptyInputs(t *testing.T) {
	result := MergeEnvs([]EnvEntry{})
	if len(result.Merged) != 0 {
		t.Errorf("expected empty merged map, got %v", result.Merged)
	}
	if len(result.Conflicts) != 0 {
		t.Errorf("expected no conflicts, got %d", len(result.Conflicts))
	}
}

func TestMergeEnvs_SameValueNoConflict(t *testing.T) {
	envs := []EnvEntry{
		{Name: "base", Values: map[string]string{"PORT": "8080"}},
		{Name: "staging", Values: map[string]string{"PORT": "8080"}},
	}
	result := MergeEnvs(envs)
	if len(result.Conflicts) != 0 {
		t.Errorf("same value should not produce conflict, got %d", len(result.Conflicts))
	}
}
