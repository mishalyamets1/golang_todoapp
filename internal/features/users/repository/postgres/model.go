package users_repository_postgres

import "github.com/mishalyamets1/golang_todoapp/internal/core/domain"

type UserModel struct {
	ID int
	Version int
	FullName string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, users := range users {
		userDomains[i] = domain.NewUser(
			users.ID,
			users.Version,
			users.FullName,
			users.PhoneNumber,
		)
	}
	return userDomains
}