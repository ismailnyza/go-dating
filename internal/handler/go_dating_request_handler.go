package handler

import (
	// "fmt"
	"sample-proj/internal/interfaces/endpoints/helloworld"
	HelloWorldService "sample-proj/services"

	httpTransport "github.com/go-kit/kit/transport/http"
)

// options customization for the Server
var options = []httpTransport.ServerOption{
    // request validation function can be written here 
    // httpTransport.ServerBefore(request.HttpToContext()),
    // httpTransport.ServerBefore(fmt.Println("request from hello world handler")),

    // encoding errors after 
    // httpTransport.ServerErrorEncoder(request.ErrorEncoder),
}

// function to handle the hello world request
func HelloWorldHandler() *httpTransport.Server{
    return httpTransport.NewServer(
        helloworld.HelloWorldEndpoint( &HelloWorldService.HelloWorldService{}), 
        helloworld.DecodeHelloWorldRequest,
        helloworld.EncodeHelloWorldResponse,
        options..., 
        )
}

// function to handle the get users request
func GetUsersHandler() *httpTransport.Server{
//     return httpTransport.NewServer(
//         getUsers.GetUsersEndpoint( &GetUsersService.GetUsersService{}),
//         )

    return httpTransport.NewServer(
        helloworld.HelloWorldEndpoint( &HelloWorldService.HelloWorldService{}), 
        helloworld.DecodeHelloWorldRequest,
        helloworld.EncodeHelloWorldResponse,
        options..., 
        )
}
