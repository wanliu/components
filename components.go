package components

import flow "github.com/trustmaster/goflow"

type GetElement struct {
	flow.Component
}

func NewGetElement() *GetElement {
	return new(GetElement)
}
