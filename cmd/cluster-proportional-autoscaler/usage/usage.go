package usage

import (
	"fmt"
)

// ensure that the list of supported modes is constant
func GetSupportedModes() []string {
	return []string{"linear", "ladder"}
}

func ValidateUsageFlag(mode string) error {
	//may eliminate this bumping to go 1.21
	errorsFound := true
	for _, availablemode := range GetSupportedModes() {
		if mode == availablemode {
			errorsFound = false
		} 
	}	
	if errorsFound {
		return fmt.Errorf("the mode selected is not supported")
	}
	return nil
}

func ShowModeUsage(mode string) {
	if mode == "linear" {
		fmt.Printf("Linear mode usage:")
	}

	if mode == "ladder" {
		fmt.Printf("Ladder mode usage:")
	}
}
