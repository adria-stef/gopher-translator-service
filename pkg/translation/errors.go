package translation

import "fmt"

type TranslationError struct {
	Cause string
}

func (e TranslationError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Cause)
}

func (e TranslationError) UserMessage() string {
	return fmt.Sprintf("provided input is invalid [%s]. please check and try again.", e.Cause)
}
