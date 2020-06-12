package controllers

import "testing"

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status got %d, want %d", got, want)
	}
}
