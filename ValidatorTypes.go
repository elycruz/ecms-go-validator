package ecms_validator

type ValueObscurator func (limit int, x string) string

type MessageFunc func (options ValidatorOptions, x interface{}) string

type MessageFuncs map[int]MessageFunc

type ValidatorOptions interface {
	GetMessageFuncs () *MessageFuncs
	GetErrorMessageByKey (key int, value interface{}) string
	GetValueObscurator () ValueObscurator
}

type ValidatorGenerator func (options ValidatorOptions) Validator

type Validator func (x interface{}) (bool, []string)
