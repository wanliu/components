package components

import flow "github.com/trustmaster/goflow"

type GetElement struct {
	flow.Component
}

func NewGetElement() interface{} {
	return new(GetElement)
}
