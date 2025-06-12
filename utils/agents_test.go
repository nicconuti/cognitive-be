package utils

import "testing"

func TestGenerateAgents(t *testing.T) {
	agents := GenerateAgents(5)
	if len(agents) != 5 {
		t.Fatalf("expected 5 agents, got %d", len(agents))
	}
	nameSet := make(map[string]bool)
	total := 0.0
	for _, a := range agents {
		if a.Score < 0 || a.Score > 1 {
			t.Errorf("score out of range: %f", a.Score)
		}
		if nameSet[a.Name] {
			t.Errorf("duplicate name detected: %s", a.Name)
		}
		nameSet[a.Name] = true
		total += a.Score
	}
	avg := total / float64(len(agents))
	if avg < 0.5 || avg > 0.8 {
		t.Errorf("unexpected mean score: %f", avg)
	}
}

func TestEstimateIQ(t *testing.T) {
	pop := GenerateAgents(20)
	iq := EstimateIQ(0.6, pop)
	if iq < 40 || iq > 160 {
		t.Errorf("iq out of expected bounds: %d", iq)
	}
}
