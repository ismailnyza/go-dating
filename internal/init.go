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
	"github.com/joho/godotenv"
)


func Init() {
    // Enable termination signal 
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    
    // Traceable context
    requestID := uuid.New().String()
    ctx := context.WithValue(context.Background(), "requestID", requestID)
    fmt.Println("server started", ctx.Value("requestID"))
    
    // init envs equal to etcd
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    } else {
        fmt.Println("DB_HOST: ", os.Getenv("DBHOST"))
    }
    // connect to db
    mainDB := NewDB("abc", "")

    goDating.InitDBConnection()
    // Initialize and store the server
    go route.HandleRequest(ctx)

    <-sigs
        route.ShutdownServer(ctx)
}
