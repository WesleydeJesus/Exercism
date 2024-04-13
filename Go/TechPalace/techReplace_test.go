package techpalace

import (
	"testing"
)

func TestWelcomeMessage(t *testing.T) {
	tests := []struct {
		name     string
		customer string
		expected string
	}{
		{"standard name", "Chris", "Welcome to the Tech Palace, CHRIS"},
		{"empty name", "", "Welcome to the Tech Palace, "},
		{"lowercase name", "alice", "Welcome to the Tech Palace, ALICE"},
		{"mixed case name", "Alexa", "Welcome to the Tech Palace, ALEXA"},
		{"uppercase name", "ALEXA", "Welcome to the Tech Palace, ALEXA"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := WelcomeMessage(test.customer)
			if got != test.expected {
				t.Errorf("WelcomeMessage(%q) = %q; want %q", test.customer, got, test.expected)
			}
		})
	}
}

func TestAddBorder(t *testing.T) {
	tests := []struct {
		name            string
		welcomeMsg      string
		numStarsPerLine int
		expected        string
	}{
		{"short message", "Welcome", 10, "**********\nWelcome\n**********"},
		{"empty message", "", 5, "*****\n\n*****"},
		{"no border", "Hello", 0, "\nHello\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := AddBorder(test.welcomeMsg, test.numStarsPerLine)
			if got != test.expected {
				t.Errorf("AddBorder(%q, %d) = %q; want %q", test.welcomeMsg, test.numStarsPerLine, got, test.expected)
			}
		})
	}
}

func TestCleanupMessage(t *testing.T) {
	tests := []struct {
		name     string
		oldMsg   string
		expected string
	}{
		{"message with stars", "****Hello****", "Hello"},
		{"message with spaces and stars", "***   Welcome home! ***", "Welcome home!"},
		{"no stars", "Just a normal message", "Just a normal message"},
		{"only stars", "*********", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CleanupMessage(test.oldMsg)
			if got != test.expected {
				t.Errorf("CleanupMessage(%q) = %q; want %q", test.oldMsg, got, test.expected)
			}
		})
	}
}