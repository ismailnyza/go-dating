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

    //enable termination signal 
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    //tracable context to maintain context throughout the application
    requestID := uuid.New().String()
    tCtx := context.WithValue(context.Background(), "requestID", requestID)


    //defer cancel the context
	ctx, cancel := context.WithCancel(tCtx)
    defer cancel()
    fmt.Println("server started" , ctx.Value("requestID"))


    //start the server
    route.HandleRequest()
    
    if <-sigs != nil {
        fmt.Println("server stopped")
    }
 
}
