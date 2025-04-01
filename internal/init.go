package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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

    // initialize 3 addapters for the application

    // db connections

    // defer a connection close 
    //garuntees that the connection is closed after the function returns  
    // defer db.Close()

    // init configs
    // devopsConfig , err := config.ReadConfig()
   
    // handle when the server has been shut down

    // initialize the routers 

    select {
    case <-sigs:
    fmt.Println("terminating the server")
     // route.StopServer(ctx)
    }
 
}
