package error

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gopi-frame/contract/validation"
	"strings"
	"text/template"
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
	code            string
	message         string
	customMessage   string
	params          []validation.Param
	translator      validation.Translator
	renderedMessage string
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
	if e.renderedMessage != "" {
		return e.renderedMessage
	}
	params := map[string]any{}
	for _, param := range e.params {
		var value = param.Value()
		if param.Key() == "attribute" {
			if e.translator != nil {
				value = e.translator.T(fmt.Sprintf("attribute.%s", param.Value()), nil)
			}
		}
		params[param.Key()] = value
	}
	var message string
	if e.customMessage != "" {
		message = e.customMessage
		goto RenderMessage
	} else {
		message = e.message
	}
	if e.translator != nil {
		e.renderedMessage = e.translator.T(e.code, params)
		return e.renderedMessage
	}
RenderMessage:
	var sb = new(strings.Builder)
	if err := template.Must(template.New("").Parse(message)).Execute(sb, params); err != nil {
		panic(err)
	}
	e.renderedMessage = sb.String()
	return e.renderedMessage
}

func (e *Error) SetMessage(message string) validation.Error {
	e.customMessage = message
	return e
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

type Bag struct {
	errors         map[string]Errors
	messages       map[string][]string
	customMessages map[string]map[string]string
}

func NewBag() *Bag {
	return &Bag{
		errors:         make(map[string]Errors),
		messages:       make(map[string][]string),
		customMessages: make(map[string]map[string]string),
	}
}

func (e *Bag) Fails() bool {
	return len(e.Failed()) > 0
}

func (e *Bag) Failed() []string {
	var keys []string
	for key := range e.errors {
		if len(e.errors[key]) > 0 {
			keys = append(keys, key)
		}
	}
	return keys
}

func (e *Bag) FailedAt(key string, codes ...string) bool {
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

func (e *Bag) HasError(key string) bool {
	if e.errors == nil {
		return false
	}
	return len(e.errors[key]) > 0
}

func (e *Bag) AddError(key string, err validation.Error) {
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

func (e *Bag) GetAllErrors() map[string]validation.Errors {
	errs := make(map[string]validation.Errors)
	for key, errorList := range e.errors {
		errs[key] = NewErrors()
		for _, err := range errorList {
			errs[key].Add(err)
		}
	}
	return errs
}

func (e *Bag) GetErrors(key string) validation.Errors {
	return e.errors[key]
}

func (e *Bag) GetError(key string, code string) validation.Error {
	return e.errors[key].Get(code)
}

// GetMessages returns the error messages for all keys.
// Once the messages are retrieved, they are stored in the messages map.
func (e *Bag) GetMessages() map[string][]string {
	var errorBag = make(map[string][]string)
	for key := range e.errors {
		errorBag[key] = e.GetMessage(key)
	}
	return errorBag
}

// SetMessages sets custom error messages
func (e *Bag) SetMessages(messages map[string]map[string]string) validation.ErrorBag {
	for key, errs := range e.errors {
		if customMessages, ok := messages[key]; ok {
			for code, err := range errs {
				if customMessage, ok := customMessages[code]; ok {
					e.errors[key][code] = err.SetMessage(customMessage)
				}
			}
		}
	}
	return e
}

// GetMessage returns the error message for the given key.
// Once the message is retrieved, it is stored in the messages map.
func (e *Bag) GetMessage(key string) []string {
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

func (e *Bag) Each(f func(key string, errs validation.Errors) bool) {
	for key, errs := range e.errors {
		if !f(key, errs) {
			break
		}
	}
}

func (e *Bag) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.errors)
}

func (e *Bag) Error() string {
	messages := e.GetMessages()
	sb := new(strings.Builder)
	for key, msgs := range messages {
		for _, msg := range msgs {
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, msg))
		}
	}
	return sb.String()
}

func (e *Bag) SetMessage(_ string) validation.Error {
	return e
}

func (e *Bag) Message() string {
	return e.Error()
}

func (e *Bag) Code() string {
	return ""
}

func (e *Bag) Params() []validation.Param {
	return nil
}

func (e *Bag) AddParam(param validation.Param) validation.Error {
	for key := range e.errors {
		for i := range e.errors[key] {
			e.errors[key][i] = e.errors[key][i].AddParam(param)
		}
	}
	return e
}
