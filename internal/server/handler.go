package server

import (
	"context"
	"fmt"
	"net/http"
	"sample-proj/internal/interfaces/endpoints/helloworld"

	// "git.mytaxi.lk/pickme/go/util/request"
	httpTransport "github.com/go-kit/kit/transport/http"
)

// options customization for the server
var options = []httpTransport.ServerOption{
    // request validation function can be written here 
    // httpTransport.ServerBefore(request.HttpToContext()),
httpTransport.ServerBefore(func(ctx context.Context, _ *http.Request) context.Context {
    fmt.Println("request from hello world handler")
    return ctx
}),

    // encoding errors after 
    // httpTransport.ServerErrorEncoder(request.ErrorEncoder),
}

func HelloWorldHandler() *httpTransport.Server{
    return httpTransport.NewServer(
        helloworld.HelloWorldEndpoint( &helloworld.HelloWorldService{}), 
        helloworld.DecodeHelloWorldRequest,
        helloworld.EncodeHelloWorldResponse,
        options..., 
        )
    }
