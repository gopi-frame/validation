# validation
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/validation.svg)](https://pkg.go.dev/github.com/gopi-frame/validation)
[![Go](https://github.com/gopi-frame/validation/actions/workflows/go.yml/badge.svg)](https://github.com/gopi-frame/validation/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/validation)](https://goreportcard.com/report/github.com/gopi-frame/validation)
[![codecov](https://codecov.io/gh/gopi-frame/validation/graph/badge.svg?token=MM0GENCM7V)](https://codecov.io/gh/gopi-frame/validation)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package validation provides a validation library for Go.

# Installation
```shell
go get -u -v github.com/gopi-frame/validation
```

# Import
```go
import "github.com/gopi-frame/validation"
```

# Usage

## Quick Start
```go
package main

import (
    "context"
    "github.com/gopi-frame/validation"
)

var user = struct {
    Name string
    Age int
}{
    Name: "gopi",
    Age: 25,
}

func main() {
    validated := validation.Validate(
        context.Background(),
        validation.NotBlank("name", user.Name),
        validation.NotBlank("age", user.Age),
        validation.GreaterThan("age", user.Age, 18),
    )
    if validated.Fails() {
        fmt.Println(validated.GetMessages())
    }    
}
```

Built-in common [validator builders](https://pkg.go.dev/github.com/gopi-frame/contract/validation#ValidatorBuilder):

- Generic type builders:
  * `validation.NotBlank[T comparable]` validates if the value is not blank
  * `validation.Blank[T comparable]` validates if the value is blank
  * `validation.In[T comparable]` validates if the value is in the given list
  * `validation.NotIn[T comparable]` validates if the value is not in the given list
  * `validation.EqualTo[T comparable]` validates if the value is equal to the given value
  * `validation.NotEqualTo[T comparable]` validates if the value is not equal to the given value
  * `validation.LessThan[T constraints.Ordered]` validates if the value is less than the given value
  * `validation.LessThanOrEqualTo[T constraints.Ordered]` validates if the value is less than or equal to the given value
  * `validation.GreaterThan[T constraints.Ordered]` validates if the value is greater than the given value
  * `validation.GreaterThanOrEqualTo[T constraints.Ordered]` validates if the value is greater than or equal to the given
    value
  
- String builders:
  * `validation.Length` validates if the length of the value is equal to the given value
  * `validation.MinLength` validates if the length of the value is greater than or equal to the given value
  * `validation.MaxLength` validates if the length of the value is less than or equal to the given value
  * `validation.StartsWith` validates if the value starts with the given value
  * `validation.StartsWithAny` validates if the value starts with any of the given values
  * `validation.EndsWith` validates if the value ends with the given value
  * `validation.EndsWithAny` validates if the value ends with any of the given values
  * `validation.NotStartsWith` validates if the value does not start with the given value
  * `validation.NotStartsWithAny` validates if the value does not start with any of the given values
  * `validation.NotEndsWith` validates if the value does not end with the given value
  * `validation.NotEndsWithAny` validates if the value does not end with any of the given values
  * `validation.Match` validates if the value matches the given regex
  * `validation.NotMatch` validates if the value does not match the given regex
  * `validation.Contains` validates if the value contains the given value
  * `validation.NotContains` validates if the value does not contain the given value
  * `validation.Upper` validates if the value is uppercase
  * `validation.Lower` validates if the value is lowercase
  * `validation.Alpha` validates if the value is alphabetic
  * `validation.AlphaNumeric` validates if the value is alphanumeric
  * `validation.AlphaDash` validates if the value is alphanumeric with dashes(_-)
  * `valication.Ascii` validates if the value only contains ascii characters
  * `validation.AsciiNumeric` validates if the value only contains ascii characters and numbers(0-9)
  * `validation.AsciiDash` validates if the value only contains ascii characters and numbers(0-9) and dashes(_-)
  * `validation.Number` validates if the value is a number
  * `validation.PositiveNumber` validates if the value is a positive number
  * `validation.NegativeNumber` validates if the value is a negative number
  * `validation.Integer` validates if the value is an integer
  * `validation.PositiveInteger` validates if the value is a positive integer
  * `validation.NegativeInteger` validates if the value is a negative integer
  * `validation.Decimal` validates if the value is a decimal
  * `validation.Binary` validates if the value is a binary
  * `validation.Octal` validates if the value is an octal
  * `validation.Hexadecimal` validates if the value is a hexadecimal
  
- Time string builders:
  * `validation.Time` validates if the value is a time in given format
  * `validation.ANSIC` validates if the value is a time in ANSIC format
  * `validation.UnixDate` validates if the value is a time in UnixDate format
  * `validation.RubyDate` validates if the value is a time in RubyDate format
  * `validation.RFC822` validates if the value is a time in RFC822 format
  * `validation.RFC822Z` validates if the value is a time in RFC822Z format
  * `validation.RFC850` validates if the value is a time in RFC850 format
  * `validation.RFC1123` validates if the value is a time in RFC1123 format
  * `validation.RFC1123Z` validates if the value is a time in RFC1123Z format
  * `validation.RFC3339` validates if the value is a time in RFC3339 format
  * `validation.RFC3339Nano` validates if the value is a time in RFC3339Nano format
  * `validation.Kitchen` validates if the value is a time in Kitchen format
  * `validation.Stamp` validates if the value is a time in Stamp format
  * `validation.StampMilli` validates if the value is a time in StampMilli format
  * `validation.StampMicro` validates if the value is a time in StampMicro format
  * `validation.StampNano` validates if the value is a time in StampNano format
  * `validation.DateTime` validates if the value is a time in DateTime format
  * `validation.DateOnly` validates if the value is a date in DateOnly format
  * `validation.TimeOnly` validates if the value is a time in TimeOnly format
  * `validation.Duration` validates if the value is a duration
  * `validation.Timezone` validates if the value is a timezone
  * `validation.Before` validates if the time string is before the given time string
  * `validation.After` validates if the time string is after the given time string
  * `validation.BeforeOrEqual` validates if the time string is before or equal to the given time string
  * `validation.AfterOrEqual` validates if the time string is after or equal to the given time string
  * `validation.BeforeTZ` validates if the time string is before the given time string in the given timezone
  * `validation.AfterTZ` validates if the time string is after the given time string in the given timezone
  * `validation.BeforeOrEqualTZ` validates if the time string is before or equal to the given time string in the given
    timezone
  * `validation.AfterOrEqualTZ` validates if the time string is after or equal to the given time string in the given

- Network string builders:
  * `validation.IP` validates if the value is a valid IP address
  * `validation.IPV4` validates if the value is a valid IPV4 address
  * `validation.IPV6` validates if the value is a valid IPV6 address
  * `validation.URL` validates if the value is a valid URL
  * `validation.URLWithScheme` validates if the value is a valid URL with given scheme
  * `validation.RequestURI` validates if the value is a valid request URI
  * `validation.URLQuery` validates if the value is a valid URL query

- Data string builders:
  * `validation.JSON` validates if the value is a valid JSON
  * `validation.JSONArray` validates if the value is a valid JSON array
  * `validation.JSONObject` validates if the value is a valid JSON object
  * `validation.JSONString` validates if the value is a valid JSON string
  * `validation.UUID` validates if the value is a valid UUID
  * `validation.UUIDv1` validates if the value is a valid version-1 UUID
  * `validation.UUIDv2` validates if the value is a valid version-2 UUID
  * `validation.UUIDv3` validates if the value is a valid version-3 UUID
  * `validation.UUIDv4` validates if the value is a valid version-4 UUID
  * `validation.UUIDv5` validates if the value is a valid version-5 UUID
  * `validation.ULID` validates if the value is a valid ULID
  * `validation.Base64` validates if the value is a valid Base64 encoded string
  * `validation.Base32` validates if the value is a valid Base32 encoded string

- Map builders:
  * `validation.ContainsKey` validates if the map contains the given key
  * `validation.ContainsValue` validates if the map contains the given value

- Slice builders:
  * `validation.Includes` validates if the slice contains the given values
  * `validation.Excludes` validates if the slice does not contain the given values
  * `validation.Unique` validates if the slice contains unique values
  * `validation.Count` validates if the slice contains the given number of values
  * `validation.MinCount` validates if the slice contains at least the given number of values
  * `validation.MaxCount` validates if the slice contains at most the given number of values

- Group builders:
  * `validation.Group[T any]` validates if the given values are valid according to the given rules
  * `validation.If[T any]` validates if the given value is valid according to the given rules when the given condition
    is true
  * `validation.Each[T any]` validates if the given values' each element is valid according to the given rules when the given
    condition

## Custom Validator

```go
package main

import (
  "context"
  "fmt"
  "github.com/gopi-frame/validation"
  "github.com/gopi-frame/validation/code"
)

func main() {
  validator := validation.NewValidator(
    validation.WithMessages(map[string]string{
      code.IsNotBlank: "{{.attribute}} is required",
    }), // replace the default messages with custom messages
  )
  // replace the default validator with the custom validator
  validation.SetDefaultValidator(validator)
  validated := validation.Validate(
    context.Background(),
    validation.NotBlank("name", "gopi"),
    validation.NotBlank("age", 25),
    validation.GreaterThan("age", 25, 18),
  )
  if validated.Fails() {
    fmt.Println(validated.GetMessages())
  }
}
```

## Set Custom Error Messages Temporarily
```go
package main

import (
    "context"
    "github.com/gopi-frame/validation"
    "github.com/gopi-frame/validation/code"
)

var user = struct {
    Name string
    Age int
}{
    Name: "gopi",
    Age: 25,
}

func main() {
    validated := validation.Validate(
        context.Background(),
        validation.NotBlank("name", user.Name),
        validation.NotBlank("age", user.Age),
        validation.GreaterThan("age", user.Age, 18),
    ).SetMessages(map[string]map[string]string{
        "name": {
            code.NotBlank: "Name is required",
        },
        "age": {
            code.NotBlank:    "Age is required",
            code.GreaterThan: "Age must be greater than 18",
        },
    })
    if validated.Fails() {
        fmt.Println(validated.GetMessages())
    }
}
```

## Conditional Validation

Conditional validation is a way to validate a value against multiple rules based on a condition.
If the value is an implementation of the `validation.Validatable` interface, the `Validate` method of the value will be
called to validate the value first.

```go
package main

import (
    "context"
    "github.com/gopi-frame/validation"
)

var user = struct {
    Name string
    Age int
    Gender string
}{
    Name: "gopi",
    Age: 25,
    Gender: "male",
}

func main() {
    var femaleOnly = true
    validated := validation.Validate(
        context.Background(),
        validation.NotBlank("name", user.Name),
        validation.NotBlank("age", user.Age),
        validation.GreaterThan("age", user.Age, 18),
        validation.If(
            femaleOnly,
            validation.EqualTo("gender", user.Gender, "female"),
        ),
    )
}
```

## Group Validation

Group validation is a way to validate a value against multiple rules.
If the value is an implementation of the `validation.Validatable` interface, the `Validate` method of the value will be
called to validate the value first.

```go
package main

import (
    "context"
    "github.com/gopi-frame/validation"
    "github.com/gopi-frame/validation/validator"
)

var user = struct {
    Name string
    Age  int
}{
    Name: "gopi",
    Age:  25,
}

func main() {
    validated := validation.Validate(
        context.Background(),
        validation.Group("Name", user.Name, validator.IsNotBlank[string]()),
        validation.Group("Age", user.Age, validator.IsNotBlank[int](), validator.IsGreaterThan[int](18)),
    )
    if validated.Fails() {
        fmt.Println(validated.GetMessages())
    }
}
```

## Custom Validatable Type

```go
package main

import (
    "context"
    vc "github.com/gopi-frame/contract/validation"
    "github.com/gopi-frame/validation"
    "github.com/gopi-frame/validation/error"
    "github.com/gopi-frame/validation/validator"
)

type User struct {
    Name string
    Age  int
}

func (u *User) Validate(ctx context.Context, _ vc.ErrorBuilder) vc.Error {
    var errs = error.NewBag()
    if validated := validation.Value(ctx, u.Name, validator.IsNotBlank[string]()); validated.Fails() {
        errs.AddError("Name", validated)
    }
    if validated := validation.Value(
        ctx, 
        u.Age, 
        validator.IsNotBlank[int](), 
        validator.IsGreaterThanOrEqualTo(18),
    ); validated.Fails() {
        errs.AddError("Age", validated)
    }
    return errs
}

func main() {
    var user = &User{
        Name: "gopi",
        Age:  25,
    }
    validated := validation.Validate(context.Context(), user)
    if validated.Fails() {
        fmt.Println(validated.GetMessages())
    }
}
```

## List Validation

```go
package main

import (
    "context"
    "github.com/gopi-frame/validation"
    "github.com/gopi-frame/validation/validator"
)


type User struct {
    Name string
    Age  int
}

func (u *User) Validate(ctx context.Context, _ vc.ErrorBuilder) vc.Error {
    var errs = error.NewBag()
    if validated := validation.Value(ctx, u.Name, validator.IsNotBlank[string]()); validated.Fails() {
        errs.AddError("Name", validated)
    }
    if validated := validation.Value(
        ctx,
        u.Age,
        validator.IsNotBlank[int](),
        validator.IsGreaterThanOrEqualTo(18),
    ); validated.Fails() {
        errs.AddError("Age", validated)
    }
    return errs
}

var users = []User{
    {"gopi", 25},
    {"john", 30},
    {"jane", 28},
}

func main() {
    validated := validation.Validate(
        context.Background(),
        validation.Each("age", users),
    )
}
```

# Translation

## Register New Translation
```go
package main

import (
  "context"
  "fmt"
  "github.com/gopi-frame/validation"
  "github.com/gopi-frame/validation/code"
  "github.com/gopi-frame/validation/translator"
)

func main() {
  translator.RegisterTranslation("zh-CN", map[string]string{
    code.IsNotBlank: "{{.attribute}}不能为空",
  })
  validated := validation.Validate(
    validation.BindLanguage(context.Background(), "zh-CN"),
    validation.NotBlank("name", ""),
    validation.GreaterThan("age", 14, 18),
  )
  if validated.Fails() {
    fmt.Println(validated.GetErrors("name").Get(code.IsNotBlank).Error()) // name不能为空
    fmt.Println(validated.GetErrors("age").Get(code.GreaterThan).Error()) // age should be greater than 18.
  }
}
```

## Register Custom Translator
Implement the [`translator.Translator`](https://pkg.go.dev/github.com/gopi-frame/contract/validation#Translator) interface and register the translator by fallowing way.
```go
validator := validation.NewValidator(
    validation.WithTranslator(new(MyTranslator)),
)
```