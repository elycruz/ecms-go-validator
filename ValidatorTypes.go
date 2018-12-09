package ecms_validator

type ValueObscurator func (limit int, x interface{}) string

type MessageTemplateFunc func (options ValidatorOptions, x interface{}) string

type MessageTemplateFuncs map[int]MessageTemplateFunc

type ValidatorOptions interface {
	GetMessageTemplates () *MessageTemplateFuncs
	GetErrorMessageByKey (key int, value interface{}) string
	GetValueObscurator () ValueObscurator
}

type ValidatorGenerator func (options ValidatorOptions) Validator

type Validator func (x interface{}) (bool, []string)
