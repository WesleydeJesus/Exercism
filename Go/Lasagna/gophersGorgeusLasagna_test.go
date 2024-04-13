package lasagna

import(
	"testing"
)

func TestRemainingOvenTime(t *testing.T) {
	tests := []struct {
		name                string
		input               int
		expectedRemaining   int
	}{
		{"No time in oven", 0, 40},
		{"Some time in oven", 10, 30},
		{"Exact time in oven", 40, 0},
		{"More than needed", 50, -10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := RemainingOvenTime(test.input); got != test.expectedRemaining {
				t.Errorf("RemainingOvenTime(%d) = %d; want %d", test.input, got, test.expectedRemaining)
			}
		})
	}
}

func TestPreparationTime(t *testing.T) {
	tests := []struct {
		name            string
		layers          int
		expectedTime    int
	}{
		{"Zero layers", 0, 0},
		{"One layer", 1, 2},
		{"Multiple layers", 5, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := PreparationTime(test.layers); got != test.expectedTime {
				t.Errorf("PreparationTime(%d) = %d; want %d", test.layers, got, test.expectedTime)
			}
		})
	}
}

func TestElapsedTime(t *testing.T) {
	tests := []struct {
		name               string
		layers             int
		minutesInOven      int
		expectedTotalTime  int
	}{
		{"No layers or oven time", 0, 0, 0},
		{"Some layers and oven time", 3, 20, 26},
		{"More layers, no oven time", 5, 0, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ElapsedTime(test.layers, test.minutesInOven); got != test.expectedTotalTime {
				t.Errorf("ElapsedTime(%d, %d) = %d; want %d", test.layers, test.minutesInOven, got, test.expectedTotalTime)
			}
		})
	}
}