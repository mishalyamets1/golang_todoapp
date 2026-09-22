package user_transport_http

import "github.com/mishalyamets1/golang_todoapp/internal/core/domain"

type UserDTOResponse struct {
	ID int `json:"id" example:"10"`
	Version int `json:"version" example:"1"`
	FullName string `json:"full_name" example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+79885530945"`
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