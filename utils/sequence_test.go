package utils

import "testing"

func TestGenerateSequenceSet(t *testing.T) {
	items := GenerateSequenceSet(3)
	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}
	for _, it := range items {
		if len(it.Series) != 4 {
			t.Errorf("series length should be 4, got %d", len(it.Series))
		}
	}
}
