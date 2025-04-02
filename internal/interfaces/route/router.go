package route

import (
	"fmt"
	"net/http"
	"sample-proj/internal/handler"
	"time"

	"github.com/gorilla/mux"
	"context"
)

var httpServer *http.Server

func HandleRequest(ctx context.Context){
    // handle the request here 
    // use the handler functions to handle the request
    router := mux.NewRouter()
    port := "8080"
    // api := router.PathPrefix("/api/v1").Subrouter()
    router.Handle("/helloworld", handler.HelloWorldHandler()).Methods(http.MethodGet)

    router.Handle("/getUsers", handler.GetUsersHandler()).Methods(http.MethodGet)
    StartServer(port, router, ctx)
}

func StartServer(p string,r http.Handler, ctx context.Context){
    serverIsRunning := make(chan bool)
    handler := r
    port := p 
    host := "127.0.0.1"

    server := &http.Server{
        Handler: handler,
        Addr:host + ":" + port, 

        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }

    go func (ctx context.Context){
        fmt.Printf("Server starting on port %s\n", port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("Error starting server: %s\n", err)
        }
        serverIsRunning <- true
        
    }(ctx)
    // annonymus functions can be executed imidiately by passing an argument at the closing bracket

<-serverIsRunning
}

// func to shutdown server
func ShutdownServer(ctx context.Context){

    if httpServer == nil {
        fmt.Println()
        fmt.Println("server Manually shutdown")
        return
    }
    if err := httpServer.Shutdown(ctx); err != nil {
        fmt.Printf("server shutdown error")
    }
        fmt.Println("server stopped")
}
