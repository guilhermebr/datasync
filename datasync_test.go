package main

import "testing"

func TestLoadProjects(t *testing.T) {
	LoadProjects()
	if len(Projects) != 1 {
		t.Fatalf("expected 1 project, got %v", len(Projects))
	}
}
