package types

//UserResponse ...
type UserResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//NewUserResponse ...
func NewUserResponse(
	email string,
	firstName string,
	lastName string) *UserResponse {
	return &UserResponse{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
