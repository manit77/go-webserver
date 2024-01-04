package main

import (
	"errors"
	"reflect"
)

type webMethod string

const (
	POST webMethod = "POST"
	GET  webMethod = "GET"
)

type webController struct {
}

func (t *webController) GetRoutes() []webRoute {

	var routes []webRoute
	var route webRoute

	route = webRoute{}
	route.authenticted = true
	route.routePath = "/"
	route.method = GET
	route.funcCall = helloWorld
	route.postDataType = reflect.TypeOf("")
	route.returnDataType = reflect.TypeOf("")
	routes = append(routes, route)

	route = webRoute{}
	route.authenticted = true
	route.routePath = "/login"
	route.method = POST
	route.funcCall = login
	route.postDataType = reflect.TypeOf((loginPost{}))
	route.returnDataType = reflect.TypeOf("")
	routes = append(routes, route)

	return routes
}

func helloWorld(postData any) (any, error) {
	return postData, nil
}

func login(postData any) (any, error) {
	var postValue = reflect.ValueOf(postData).Elem().Interface().(loginPost)
	var result loginResult

	if postValue.Username == "a" && postValue.Password == "b" {
		result.AuthToken = "here is your authtoken"
		return result, nil
	} else {
		return nil, errors.New("invalid login")
	}

}
