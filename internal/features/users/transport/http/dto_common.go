package user_transport_http

import "github.com/mishalyamets1/golang_todoapp/internal/core/domain"

type UserDTOResponse struct {
	ID int `json:"id"`
	Version int `json:"version"`
	FullName string `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID: user.ID,
		Version: user.Version,
		FullName: user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains (users []domain.User) []UserDTOResponse {
	usersDto := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDto[i] = userDTOFromDomain(user)
	}
	return usersDto
}