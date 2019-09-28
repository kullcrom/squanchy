package types

//User ...
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUser() *User {
	return &User{
		ID:        0,
		Email:     "email",
		FirstName: "firstname",
		LastName:  "lastname",
	}
}
