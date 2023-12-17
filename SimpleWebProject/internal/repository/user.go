package repository

import "SimpleWebProject/internal/model"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetAll() (users []*model.User) {

	users = []*model.User{
		{
			ID:        1,
			FirstName: "Marsel",
			LastName:  "Sidikov",
		}, {
			ID:        2,
			FirstName: "Maxim",
			LastName:  "Pozdeev",
		},
	}

	return
}
