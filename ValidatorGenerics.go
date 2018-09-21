package ecms_validator

// GetErrorMessageByKey is a generic function that gets the message function from a collection
// of message template functions and calls it using passed in options and value values.
func GetErrorMessageByKey (options ValidatorOptions, key int, value interface{}) string {
	messageTemplates := options.GetMessageTemplates()
	if messageTemplates[key] != nil {
		return messageTemplates[key](options, value)
	}
	return "No error message found."
}

func DefaultValueObscurator (x interface{}) string {
	return ""
}
