package main

import "github.com/wanliu/components"

var _components = make(map[string]func() interface{})

func ComponentList() map[string]func() interface{} {
	_components["dom/GetElement"] = (func() interface{})(components.NewGetElement)
	return _components
}
