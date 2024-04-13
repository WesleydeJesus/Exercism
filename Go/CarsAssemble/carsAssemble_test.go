package cars

import (
	"testing"
)

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	tests := []struct {
		productionRate int
		successRate    float64
		expected       float64
	}{
		{100, 100.0, 100.0},
		{100, 80.0, 80.0},
		{0, 100.0, 0.0},
		{200, 100.0, 200.0},
		{200, 0.0, 0.0},
		{200, 50.0, 100.0},
		{123, 78.5, 96.555},
		{1000, 33.333, 333.33},
	}

	for _, test := range tests {
		if got := CalculateWorkingCarsPerHour(test.productionRate, test.successRate); got != test.expected {
			t.Errorf("CalculateWorkingCarsPerHour(%d, %f) = %f; want %f", test.productionRate, test.successRate, got, test.expected)
		}
	}
}

func TestCalculateWorkingCarsPerMinute(t *testing.T) {
	tests := []struct {
		productionRate int
		successRate    float64
		expected       int
	}{
		{100, 100.0, 1},
		{100, 80.0, 1},
		{0, 100.0, 0},
		{100, 50.0, 0},
		{60, 100.0, 1},
		{350, 100.0, 5},
		{350, 40.0, 2},
		{999, 1.0, 0},
	}

	for _, test := range tests {
		if got := CalculateWorkingCarsPerMinute(test.productionRate, test.successRate); got != test.expected {
			t.Errorf("CalculateWorkingCarsPerMinute(%d, %f) = %d; want %d", test.productionRate, test.successRate, got, test.expected)
		}
	}
}

func TestCalculateCost(t *testing.T) {
	tests := []struct {
		carsCount int
		expected  uint
	}{
		{1, 10000},
		{10, 95000},
		{11, 105000},
		{0, 0},
		{9, 90000},
		{20, 190000},
		{21, 200000},
		{29, 280000},
		{99, 945000},
		{100, 950000},
		{101, 960000},
	}

	for _, test := range tests {
		if got := CalculateCost(test.carsCount); got != test.expected {
			t.Errorf("CalculateCost(%d) = %d; want %d", test.carsCount, got, test.expected)
		}
	}
}
