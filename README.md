# ecms-go-validator
Validator package inspired by zend-validator.

Also (I know) there are good validator packages already available for go (for example, go-ozzo/ozzo-validator comes close to zend framework's validator 
though merges the interface as part of the package), 
though for some needs lower level interfaces are required.

For the zend framework version of the `Validator`, `Input` and `InputFilter` classes the `Validator` class
is a pluggable unit.  The validation parts, in an application, usually happen from the `Input` and `InputFilter` classes.
This library is just the `Validator`'s part.  

## Common Use Cases
- For creating input validation classes (constructors that can have multiple validators and/or filters for validating and filtering a given field in a dataset).
- For creating input filter chain classes (constructors that take multiple inputs 
and validates them (inputs have to be instances of some input validation class)) return a conglomerate of results as one result.
- @note ZendFramework/Validator ZendFramework/InputFilter do exactly what is mentioned
in the above bullet points. 
 
## Docs
@todo Add docs

## Examples:

Given our default types:
```go
// Default types
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

A trivial validator would look something like:
```go
package somepackage

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

func NewNotNilValidatorOptions () NotNilValidatorOptions {
	return NotNilValidatorOptions{
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

func (n NotNilValidatorOptions) GetErrorMessageByKey (key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n NotNilValidatorOptions) GetMessageFuncs () *MessageFuncs {
	return n.MessageFuncs
}

func (n NotNilValidatorOptions) GetValueObscurator () ValueObscurator {
	return DefaultValueObscurator
}

```

## Mvp Todos
- [x] - Remove `ValidationResult` struct.  We ca return multiple values in go.  There is no need for `ValidationResult`.

## Tentative Todos
- [X] - ~~Change Validator signature to `func (x interface{}) []string` (requires more evaluation/consensus).~~  We're going to keep the current return format for validators;  I.e., `(bool, []string{})`;

## License
BSD-3-Clause
