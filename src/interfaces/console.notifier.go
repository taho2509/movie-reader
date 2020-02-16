package interfaces

import (
	"fmt"
)

type consoleNotifier struct{}

func (*consoleNotifier) Notify(title string) {
	fmt.Println("movie: ", title)
}

// NewConsoleNotifier return a concrete implementation of a DataNotifier
func NewConsoleNotifier() *consoleNotifier {
	return &consoleNotifier{}
}
