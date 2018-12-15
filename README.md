# ecms-go-validator
Validator package inspired by zend-validator.

Also (I know) there are good validator packages already available,
for go in the wild, though none take the Zend Framework approach to validators (though go-ozzo/ozzo-validation comes pretty close);  I.e.,

- Given some options a validator can be generated from said options.
- Given options can be of the *-validator-options type (each validator generator can 
have it's own options type).
- Options must contain a property which contains a map 
of error message Ids to possible error messages and/or functions that generate error messages.

for example:

Types can look like:
```go
package ecms_validator

// For obscuring values in error messages (composed into validators 
//  where required (for example in a credit-card validator)) 
type ValueObscurator func (limit int, x interface{}) string

type MessageFunc func (options ValidatorOptions, x interface{}) string

type MessageFuncs map[int]MessageFunc

type ValidatorOptions interface {
	GetMessageFuncs () *MessageFuncs
	GetErrorMessageByKey (key int, value interface{}) string
	GetValueObscurator () ValueObscurator
}

type ValidatorGenerator func (options ValidatorOptions) Validator

type Validator func (x interface{}) (bool, []string)

```

And an example validator can look like:
```go
package ecms_validator

import (
	. "github.com/extensible-cms/ecms-go-validator"
	. "github.com/extensible-cms/ecms-go-validator/is"
)

type NotNilValidatorOptions struct {
	MessageFuncs *MessageFuncs
}

const (
	DefaultNotNilMsg = "`nil` is not allowed as a value."
	NilNotAllowed = 0x0000ff
)

func NewNotEmptyValidatorOptions () NotEmptyValidatorOptions {
	return NotEmptyValidatorOptions{
		MessageFuncs: &MessageFuncs{
			NilNotAllowed: func(ops ValidatorOptions, x interface{}) string {
				return DefaultNotNilMsg
			},
		},
	}
}

func NotNilValidator (options ValidatorOptions) Validator {
	return func (x interface{}) (bool, []string){
		if x == nil {
			return false, []string{GetErrorMessageByKey(options, NilNotAllowed, x)}
		}
		return true, nil
	}
}
```

## Mvp Todos
- [x] - Remove `ValidationResult` struct.  We ca return multiple values in go.  There is no need for `ValidationResult`.

## Tentative Todos
- [ ] - Change Validator signature to `func (x interface{}) []string` (requires more evaluation/consensus).

## License
BSD-3-Clause
