package test

import "testing"

func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected: %v; actual: %v", expected, actual)
	}
}

func MapContains[K comparable, V any](t *testing.T, m map[K]V, keys ...K) {
	t.Helper()

	for _, key := range keys {
		_, found := m[key]
		if !found {
			t.Errorf("expected key: %v; to be in map: %v", key, m)
		}
	}
}
