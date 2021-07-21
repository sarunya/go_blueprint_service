package base

import (
	"bp_service/lib/structs"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartServer(dependencies structs.Depend, routes []structs.Route) {

	router := httprouter.New()
	for _, route := range routes {
		fmt.Println(route)
		router.Handle(string(route.Method), route.Path, route.Handler)
	}
	http.ListenAndServe(":8080", router)

}

func wrapHandler(route structs.Route) {
	//schema validation

	// call route handler

	//response to be sent to client
}
