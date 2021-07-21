package routes

import (
	"bp_service/lib/structs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetRoutes() []structs.Route {
	publicRoutes := []structs.Route{
		{
			Path:   "/task/create",
			Method: structs.Method(structs.POST),
			Handler: func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				w.Write([]byte("Successfully written"))
			},
			Schema: nil,
		},
	}

	return publicRoutes
}
