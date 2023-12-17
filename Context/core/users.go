package core

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}
