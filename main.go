package main

import (
	"fmt"
	"sample-proj/internal"
)

func main() {
   fmt.Printf("server started\n") 

    port := ":8080"
    fmt.Printf("Server starting on port %s\n", port)
    internal.Init()

}
