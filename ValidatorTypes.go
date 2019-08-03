package ecms_validator

// For obscuring values in error messages (composed into validators
//  where required (for example in a credit-card validator))
type ValueObscurator func (limit int, str string) string

type MessageFunc func (options ValidatorOptions, x interface{}) string // returns message

type MessageFuncs map[int]MessageFunc // message templates

type ValidatorOptions interface {
	GetMessageFuncs () *MessageFuncs
	GetErrorMessageByKey (key int, value interface{}) string
	GetValueObscurator () ValueObscurator
}

type ValidatorGenerator func (options ValidatorOptions) Validator

type Validator func (x interface{}) (bool, []string)
