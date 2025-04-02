package domain

import "context"

// similar to the consts file stores domain specific constants and interfaces

// implement the service interface
// 
type Service interface {
    Do(ctx context.Context, arg []interface{}) ([]interface{}, error)
}
