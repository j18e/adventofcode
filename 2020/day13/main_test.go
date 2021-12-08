package main

import (
	"testing"
)

func Test_part2(t *testing.T) {
	tests := []struct {
		steps StepSet
		exp   int
	}{
		{StepSet{0: 7, 1: 13, 4: 59, 6: 31, 7: 19}, 1068781},
		{StepSet{0: 17, 2: 13, 3: 19}, 3417},
		{StepSet{0: 67, 1: 7, 2: 59, 3: 61}, 754018},
		{StepSet{0: 67, 2: 7, 3: 59, 4: 61}, 779210},
		{StepSet{0: 67, 1: 7, 3: 59, 4: 61}, 1261476},
		{StepSet{0: 1789, 1: 37, 2: 47, 3: 1889}, 1202161486},
	}
	for i, tt := range tests {
		if got := tt.steps.Multiplier(); got != tt.exp {
			t.Errorf("case %d: got %d, expected %d", i, got, tt.exp)
		}
	}
}
