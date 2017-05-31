package components

import flow "github.com/trustmaster/goflow"

type GetElement struct {
	flow.Component
}

type Split struct {
	flow.Component
}

func NewGetElement() interface{} {
	return new(GetElement)
}

func NewSplit() interface{} {
	return new(Split)
}
