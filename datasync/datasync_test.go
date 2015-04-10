package main

import "testing"

func TestLoadProjects(t *testing.T) {
	LoadProjects()
	if len(Projects) != 2 {
		t.Fatalf("expected 2 projects, got %v", len(Projects))
	}
}
