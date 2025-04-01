package helloworld

import (
	"context"
	"encoding/json"
	"net/http"
)

func  EncodeHelloWorldResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {

    r := response.(string)
    //creating a response object (json)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    //other options http ResponseWriter can do
    // compression of the response and etc

    return json.NewEncoder(w).Encode(r)
} 
