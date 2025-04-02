package users

import "context"

// User struct
func GetAllUsersEndpoint(getAllUsersService domain.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {

        r, err := getAllUsersService.Do(ctx, []interface{}{request})

        return r[0], err
    }
    // handle the request

}
