# ecms-go-validator
Validator package inspired by zend-validator.

Also (I know) there are good validator packages already available
for go in the wild though none take the approach of Validators, Inputs, and Input-filters as units of code - This pattern offers a more composable/customizable approach;  I.e.,

Given that a validator can be generated given some options:
- An Input can have multiple validators.
- An InputFilter can have multiple Inputs.

This makes reusing inputs and validators fairly easy for given cases (validating domain specific data objects for input into some IO, for example).

Also, one more unique change to validators;  They receive, in their options struct, `MessageTemplateFuncs map[int]MessageTemplateFunc`.  An idea that allows your validators to return more than one error message for a given validation call also this idea
allows users to customize any one of the messages contained in the `MessageTemplateFuncs` map (yes all it is, is a `map[int]MessageTemplateFunc` (look at types file in src)).

## Mvp Todos
- [ ] - Remove `ValidationResult` struct.  We ca return multiple values in go.  Is there really a need for `ValidationResult`?
- [ ] - Change Validator signature to `func (x interface{}) []string`

## License
BSD-3-Clause
