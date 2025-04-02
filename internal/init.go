package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sample-proj/internal/interfaces/route"

	// 	"sample-proj/config"
	"syscall"

	//	"git.mytaxi.lk/pickme/go/util/router"
	"github.com/google/uuid"
)


func Init() {
    // Enable termination signal 
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    
    // Traceable context
    requestID := uuid.New().String()
    ctx := context.WithValue(context.Background(), "requestID", requestID)
    fmt.Println("server started", ctx.Value("requestID"))
    
    // Initialize and store the server
    go route.HandleRequest(ctx)

    <-sigs
        route.ShutdownServer(ctx)
}
