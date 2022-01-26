package class

import "github.com/goal-web/contracts"

type TypeException struct {
	error
	fields contracts.Fields
}

func (this TypeException) Error() string {
	return this.error.Error()
}

func (this TypeException) Fields() contracts.Fields {
	return this.fields
}
