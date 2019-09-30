package types

//User ...
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//NewUser ...
func NewUser(
	id int,
	email string,
	password string,
	firstName string,
	lastName string) *User {
	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
}
