package ecms_validator

// GetErrorMessageByKey is a generic function that gets the message function from a collection
// of message template functions and calls it using passed in options and value values.
func GetErrorMessageByKey (options ValidatorOptions, key int, value interface{}) string {
	MessageFuncs := *options.GetMessageFuncs()
	if MessageFuncs[key] != nil {
		return MessageFuncs[key](options, value)
	}
	return "No error message found."
}

// DefaultValueObscurator Returns an obscured string representation of given
// value.  Obscures string using "*" character up to given `limit`.
func DefaultValueObscurator (limit int, str string) string {
	return ObscurateLeft(limit, str)
}

func ObscurateLeft (numChars int, str string) string {
	xLen := len(str)
	if xLen == 0 || numChars <= 0 {
		return str
	}
	char := "*"
	out := ""
	for i := 0; i < numChars; i += 1 {
		out += char
	}
	return out + str[numChars:xLen]
}

func ObscurateRight(numChars int, str string) string  {
	xLen := len(str)
	if xLen == 0 || numChars <= 0 {
		return str
	}
	char := "*"
	out := ""
	for i := 0; i < numChars; i += 1 {
		out += char
	}
	return str[0:xLen - numChars] + out
}