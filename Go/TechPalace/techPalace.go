package techpalace

import (
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starts := strings.Repeat("*", numStarsPerLine)
	return starts + "\n" + welcomeMsg + "\n" + starts
}

/*
reading doc is sometimes valuable:

	https://pkg.go.dev/strings#ReplaceAll
	https://pkg.go.dev/strings#Fields
	https://pkg.go.dev/strings#Join

it turned out completely illegible, but it was what I could think of
*/
func CleanupMessage(oldMsg string) string {
	return strings.Join(strings.Fields(strings.ReplaceAll(oldMsg, "*", "")), " ")
}
