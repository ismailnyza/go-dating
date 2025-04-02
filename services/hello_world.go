package HelloWorldService

import "context"

type HelloWorldService struct {}

//template of the ping service taks the context and a slide of an interface returns and interface and error 
// interface in go is eqivalent to a generic type in other languages ie any

func (h *HelloWorldService) Do(ctx context.Context, arg []interface{}) ([]interface{}, error) {
    return []interface{}{"Hello World"}, nil
}
