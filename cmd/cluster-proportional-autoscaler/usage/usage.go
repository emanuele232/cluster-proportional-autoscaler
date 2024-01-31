package usage

import (
	"fmt"
	"slices"
)

// ensure that the list of supported modes is constant
func GetSupportedModes() []string {
	return []string{"linear", "ladder"}
}

func ValidateUsageFlag(mode string) {
	if slices.Contains() 
}

func ShowModeUsage(mode string) {
	if mode == "linear" {
		fmt.Printf("Linear mode usage")
	}

	if mode == "ladder" {
		fmt.Printf("Ladder mode usage")
	}
}
