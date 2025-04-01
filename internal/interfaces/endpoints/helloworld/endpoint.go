package helloworld

import (
	"context"
	"sample-proj/domain"

	"github.com/go-kit/kit/endpoint"
)

func HelloWorldEndpoint(helloWorldService domain.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    
        r, err := helloWorldService.Do(ctx , []interface{}{request})

        return r[0], err
    }
    // handle the request
}
