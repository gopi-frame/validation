package is

import (
	"math/big"
)

func Number(s string) bool {
	_, _, err := new(big.Float).Parse(s, 0)
	return err == nil
}

func PositiveNumber(s string) bool {
	f, _, err := new(big.Float).Parse(s, 0)
	return err == nil && !f.Signbit()
}

func NegativeNumber(s string) bool {
	f, _, err := new(big.Float).Parse(s, 0)
	return err == nil && f.Signbit()
}

func Integer(s string) bool {
	f, _, err := new(big.Float).Parse(s, 0)
	return err == nil && f.IsInt()
}

func PositiveInteger(s string) bool {
	f, _, err := new(big.Float).Parse(s, 0)
	return err == nil && f.IsInt() && !f.Signbit()
}

func NegativeInteger(s string) bool {
	f, _, err := new(big.Float).Parse(s, 0)
	return err == nil && f.IsInt() && f.Signbit()
}

func Binary(s string) bool {
	var sign uint8
	if s[0] == '-' || s[0] == '+' {
		sign = s[0]
		s = s[1:]
	}
	if s[0] == '0' && (s[1] == 'b' || s[1] == 'B') {
		s = s[2:]
	}
	if sign == '-' || sign == '+' {
		s = string(sign) + s
	}
	_, _, err := new(big.Float).Parse(s, 2)
	return err == nil
}

func Octal(s string) bool {
	var sign uint8
	if s[0] == '-' || s[0] == '+' {
		sign = s[0]
		s = s[1:]
	}
	if s[0] == '0' && (s[1] == 'o' || s[1] == 'O') {
		s = s[2:]
	}
	if sign == '-' || sign == '+' {
		s = string(sign) + s
	}
	_, _, err := new(big.Float).Parse(s, 8)
	return err == nil
}

func Hexadecimal(s string) bool {
	var sign uint8
	if s[0] == '-' || s[0] == '+' {
		sign = s[0]
		s = s[1:]
	}
	if s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	if sign == '-' || sign == '+' {
		s = string(sign) + s
	}
	_, _, err := new(big.Float).Parse(s, 16)
	return err == nil
}

func Decimal(s string) bool {
	_, _, err := new(big.Float).Parse(s, 10)
	return err == nil
}
