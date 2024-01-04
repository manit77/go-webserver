package main

import (
	"reflect"
)

type webRoute struct {
	routePath      string
	authenticted   bool
	method         webMethod
	postDataType   reflect.Type
	returnDataType reflect.Type
	funcCall       func(postData any) (any, error)
}
