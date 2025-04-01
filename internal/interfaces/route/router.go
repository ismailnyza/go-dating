package route

import (
	"fmt"
	"net/http"
	"sample-proj/internal/handler"

	"github.com/gorilla/mux"
)

func HandleRequest(){
    // handle the request here 
    // use the handler functions to handle the request
    router := mux.NewRouter()
    port := ":8080"
    // api := router.PathPrefix("/api/v1").Subrouter()
    router.Handle("/helloworld", handler.HelloWorldHandler()).Methods(http.MethodGet)


err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
