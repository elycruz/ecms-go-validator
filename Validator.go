package ecms_validator

import "fmt"

// GetErrorMessageByKey is a generic function that gets the message function from a collection
// of message template functions and calls it using passed in options and value values.
func GetErrorMessageByKey (options ValidatorOptions, key int, value interface{}) string {
	MessageFuncs := *options.GetMessageFuncs()
	if MessageFuncs[key] != nil {
		return MessageFuncs[key](options, value)
	}
	return "No error message found."
}

// DefaultValueObscurator Returns and obscured string representation of given
// value.  Obscures string rep using "*" character up to given `limit`.
func DefaultValueObscurator (limit int, x interface{}) string {
	str := fmt.Sprintf("%v", x)[limit - 1:]
	char := "*"
	out := ""
	for ind := 0; ind < limit; ind += 1 {
		out += char
	}
	return out + str
}
