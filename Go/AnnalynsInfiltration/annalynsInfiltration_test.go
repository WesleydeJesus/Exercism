package annalyn

import (
	"testing"
)

func TestCanFastAttack(t *testing.T) {
	if CanFastAttack(true) {
		t.Errorf("CanFastAttack(true) = true; want false")
	}
	if !CanFastAttack(false) {
		t.Errorf("CanFastAttack(false) = false; want true")
	}
}

func TestCanSpy(t *testing.T) {
	tests := []struct {
		name       string
		knightAwake bool
		archerAwake bool
		prisonerAwake bool
		expected   bool
	}{
		{"All asleep", false, false, false, false},
		{"Knight awake", true, false, false, true},
		{"Archer awake", false, true, false, true},
		{"Prisoner awake", false, false, true, true},
		{"All awake", true, true, true, true},
	}

	for _, tt := range tests {
		if got := CanSpy(tt.knightAwake, tt.archerAwake, tt.prisonerAwake); got != tt.expected {
			t.Errorf("CanSpy(%t, %t, %t) = %t; want %t", tt.knightAwake, tt.archerAwake, tt.prisonerAwake, got, tt.expected)
		}
	}
}

func TestCanSignalPrisoner(t *testing.T) {
	if !CanSignalPrisoner(false, true) {
		t.Errorf("CanSignalPrisoner(false, true) = false; want true")
	}
	if CanSignalPrisoner(true, true) {
		t.Errorf("CanSignalPrisoner(true, true) = true; want false")
	}
	if CanSignalPrisoner(false, false) {
		t.Errorf("CanSignalPrisoner(false, false) = true; want false")
	}
	if CanSignalPrisoner(true, false) {
		t.Errorf("CanSignalPrisoner(true, false) = true; want false")
	}
}

func TestCanFreePrisoner(t *testing.T) {
	tests := []struct {
		name             string
		knightAwake      bool
		archerAwake      bool
		prisonerAwake    bool
		petDogPresent    bool
		expected         bool
	}{
		{"All asleep, no dog", false, false, true, false, true},
		{"All asleep, with dog", false, false, true, true, true},
		{"Archer awake, with dog", false, true, true, true, false},
		{"Archer awake, no dog", false, true, true, false, false},
		{"Knight and Archer awake, with dog", true, true, true, true, false},
		{"Knight and Archer awake, no dog", true, true, true, false, false},
		{"Everyone awake, with dog", true, true, true, true, false},
	}

	for _, tt := range tests {
		if got := CanFreePrisoner(tt.knightAwake, tt.archerAwake, tt.prisonerAwake, tt.petDogPresent); got != tt.expected {
			t.Errorf("%s = %t; want %t", tt.name, got, tt.expected)
		}
	}
}
