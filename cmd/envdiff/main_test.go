package main

import (
	"testing"
)

func TestEnvName_WithSuffix(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{".env.production", "production"},
		{".env.staging", "staging"},
		{".env.development", "development"},
		{".env", ".env"},
		{"configs/.env.test", "test"},
		{"configs/.env", ".env"},
		{"/absolute/path/.env.qa", "qa"},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := envName(tc.input)
			if got != tc.want {
				t.Errorf("envName(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestEnvName_NoPrefix(t *testing.T) {
	got := envName("myconfig.env")
	if got != "myconfig.env" {
		t.Errorf("expected myconfig.env, got %q", got)
	}
}

func TestEnvName_DeepPath(t *testing.T) {
	got := envName("/home/user/project/envs/.env.ci")
	want := "ci"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
