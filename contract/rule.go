package contract

type Rule interface {
	Validate(Form) bool
}
