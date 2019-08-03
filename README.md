# ecms-go-validator
A validator package which offers the parts required for validation as separate parts: messaging templates, validators, 
validator options, and validator getters.

There are already good validation libraries out there but none that separate out the components of validation into lower
level primitives (which are useful for composing your validation components in more precise manners;  Example: 
Validating with locale based content in messaging and allowing multiple messages per input).  

Inspired by zend-validator.

## Common Use Cases
- Validating an input (an input could have many validation points (length, notEmpty, specificChars etc.))
- Validating all inputs in an input map (ex: `map[string]Input{}`);  Same as "Validating form inputs".
 
## Docs
Look at `./ValidatorTypes.go` then other sources for now.
@todo Add method docs to readme and standalone.

## Examples:

Given our default types:
```go
// ./ValidatorTypes.go
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

Then usage of created validator:
```go
package somepackage
import "net/http"

type JsonModel struct {
	Request map[string]interface{}
	Errors  map[int]interface{}
	Data    map[string]interface{}
}

func NewJsonModel() JsonModel {
	return JsonModel{
		map[string]interface{}{},
		map[int]interface{}{},
		map[string]interface{}{},
	}
}

// Would preferably be instantiated with action call (and possibly seeded with some locale based error messages
//  for it's given cases)
var notNilValidator = NotNilValidator(NotNilValidatorOptions())

func SomeAction (r *http.Request, w http.ResponseWriter) {
    // Outgoing json view model
    out := NewJsonModel()
    out.Data["success"] = false

    // Or some other model
    incoming := NewJsonModel()
    
    // Ensure there are no errors in http "POST" form data
    if err := r.ParseForm(); err != nil {
        // ...
        return
    }

	// Parse incoming body
	bytesFromBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// ...
		return
	}

	// Unmarshal into struct
	if err := json.Unmarshal(bytesFromBody, &incoming); err != nil {
		// ...
		return
	}

    result, messages := notNilValidator(incoming.Data)
    if result != true {
        // ... Write header
        // ... Send back result and messages
        // Hypothetical error-codes package
        out.Errors[errorcodes.INVALID_INPUT_DATA] = errorcodes.messagesMap[errorcodes.INVALID_INPUT_DATA]
        out.Data["messages"] = messages // ui pertinent messages
        out.Data["result"] = result     // ""
        return
    }

    // Else continue with processing
}
```

@note Input and InputFilter examples coming soon (@todo).

## License
BSD-3-Clause
