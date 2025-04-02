package Services

// import ("context")

type User struct {
    UserID    string `json:"userId"`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    Height    string `json:"height"`
}

// get users Services
// func (u *User) Do(ctx context.Context, arg []interface{}) ([]interface{}, error) {
//     // user := u.
//     return nil
// }
