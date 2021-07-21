package structs

import (
	"bp_service/lib/helper"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Path string
	Method
	Schema  helper.Schema
	Handler httprouter.Handle
}

type Method string

const (
	POST   Method = "POST"
	GET    Method = "GET"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)
