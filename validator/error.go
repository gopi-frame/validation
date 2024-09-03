package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gopi-frame/contract/validation"
	"strings"
)

type ErrorParam struct {
	key   string
	value string
}

func (e *ErrorParam) Key() string {
	return e.key
}

func (e *ErrorParam) Value() string {
	return e.value
}

func NewParam(key string, value any) *ErrorParam {
	return &ErrorParam{
		key:   key,
		value: fmt.Sprintf("%v", value),
	}
}

type Error struct {
	code       string
	message    string
	params     []validation.Param
	translator validation.Translator
}

func NewError(code string, message string, params ...validation.Param) *Error {
	return &Error{
		code:    code,
		message: message,
		params:  params,
	}
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Error() string {
	params := map[string]any{}
	for _, param := range e.params {
		if param.Key() == "attribute" {
			params[param.Key()] = e.translator.T(fmt.Sprintf("attribute.%s", param.Value()), nil)
		} else {
			params[param.Key()] = param.Value()
		}
	}
	return e.translator.T(e.code, params)
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Params() []validation.Param {
	return e.params
}

func (e *Error) AddParam(param validation.Param) validation.Error {
	for i := range e.params {
		if e.params[i].Key() == param.Key() {
			e.params[i] = param
			return e
		}
	}
	e.params = append(e.params, param)
	return e
}

func (e *Error) SetTranslator(translator validation.Translator) validation.Error {
	e.translator = translator
	return e
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.Code(),
		Message: e.Error(),
	})
}

type Errors map[string]validation.Error

func NewErrors() Errors {
	return make(map[string]validation.Error)
}

func (e Errors) Error() string {
	errs := make([]string, 0, len(e))
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, ", ")
}

func (e Errors) Add(err validation.Error) {
	if _, ok := e[err.Code()]; ok {
		return
	} else {
		e[err.Code()] = err
	}
}

func (e Errors) Get(code string) validation.Error {
	if e == nil {
		return nil
	}
	return e[code]
}

func (e Errors) Has(code string) bool {
	if e == nil {
		return false
	}
	_, ok := e[code]
	return ok
}

func (e Errors) Each(f func(key string, err validation.Error) bool) {
	for code, err := range e {
		if !f(code, err) {
			break
		}
	}
}

func (e Errors) MarshalJSON() ([]byte, error) {
	var errs = make([]validation.Error, 0, len(e))
	for _, err := range e {
		errs = append(errs, err)
	}
	return json.Marshal(errs)
}

type ErrorBag struct {
	errors     map[string]Errors
	translator validation.Translator
	messages   map[string][]string
}

func NewErrorBag(translator validation.Translator) *ErrorBag {
	return &ErrorBag{
		errors:     make(map[string]Errors),
		translator: translator,
	}
}

func (e *ErrorBag) Fails() bool {
	return len(e.Failed()) > 0
}

func (e *ErrorBag) Failed() []string {
	var keys []string
	for key := range e.errors {
		if len(e.errors[key]) > 0 {
			keys = append(keys, key)
		}
	}
	return keys
}

func (e *ErrorBag) FailedAt(key string, codes ...string) bool {
	if e.errors == nil {
		return false
	}
	if len(codes) == 0 {
		return len(e.errors[key]) > 0
	}
	for _, code := range codes {
		if ok := e.errors[key].Has(code); ok {
			return true
		}
	}
	return false
}

func (e *ErrorBag) HasError(key string) bool {
	if e.errors == nil {
		return false
	}
	return len(e.errors[key]) > 0
}

func (e *ErrorBag) AddError(key string, err validation.Error) {
	if e.errors == nil {
		e.errors = make(map[string]Errors)
	}

	var errorBag validation.ErrorBag
	if errors.As(err, &errorBag) {
		errorBag.Each(func(k string, errs validation.Errors) bool {
			k0 := key
			if k != "" {
				k0 = strings.Trim(fmt.Sprintf("%s.%s", key, k), ".")
			}
			errs.Each(func(code string, err validation.Error) bool {
				e.AddError(k0, err)
				return true
			})
			return true
		})
	} else {
		if e.errors[key] == nil {
			e.errors[key] = NewErrors()
		}
		e.errors[key].Add(err)
	}
}

func (e *ErrorBag) GetAllErrors() map[string]validation.Errors {
	errs := make(map[string]validation.Errors)
	for key, errorList := range e.errors {
		errs[key] = NewErrors()
		for _, err := range errorList {
			errs[key].Add(err)
		}
	}
	return errs
}

func (e *ErrorBag) GetErrors(key string) validation.Errors {
	return e.errors[key]
}

func (e *ErrorBag) GetError(key string, code string) validation.Error {
	return e.errors[key].Get(code)
}

// GetMessages returns the error messages for all keys.
// Once the messages are retrieved, they are stored in the messages map.
func (e *ErrorBag) GetMessages() map[string][]string {
	var errorBag = make(map[string][]string)
	for key := range e.errors {
		errorBag[key] = e.GetMessage(key)
	}
	return errorBag
}

// GetMessage returns the error message for the given key.
// Once the message is retrieved, it is stored in the messages map.
func (e *ErrorBag) GetMessage(key string) []string {
	if e.messages == nil {
		e.messages = make(map[string][]string)
	}
	if messages, ok := e.messages[key]; ok {
		return messages
	}
	var errorList []string
	for _, err := range e.errors[key] {
		errorList = append(errorList, err.Error())
	}
	e.messages[key] = errorList
	return errorList
}

func (e *ErrorBag) Each(f func(key string, errs validation.Errors) bool) {
	for key, errs := range e.errors {
		if !f(key, errs) {
			break
		}
	}
}

func (e *ErrorBag) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.errors)
}

func (e *ErrorBag) Error() string {
	messages := e.GetMessages()
	sb := new(strings.Builder)
	for key, msgs := range messages {
		for _, msg := range msgs {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, msg))
		}
	}
	return sb.String()
}

func (e *ErrorBag) Message() string {
	return e.Error()
}

func (e *ErrorBag) Code() string {
	return ""
}

func (e *ErrorBag) Params() []validation.Param {
	return nil
}

func (e *ErrorBag) AddParam(param validation.Param) validation.Error {
	for key := range e.errors {
		for i := range e.errors[key] {
			e.errors[key][i] = e.errors[key][i].AddParam(param)
		}
	}
	return e
}
