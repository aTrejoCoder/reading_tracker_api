package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
)

type UserMapper struct {
}

func (um UserMapper) EntityToDTO(user models.User) dtos.UserDTO {
	return dtos.UserDTO{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (um UserMapper) InsertDtoToEntity(userInsertDTO dtos.UserInsertDTO) models.User {
	now := time.Now().UTC().Local()
	return models.User{
		Username:  userInsertDTO.Username,
		Email:     userInsertDTO.Email,
		Password:  userInsertDTO.Password,
		CreatedAt: now,
		UpdatedAt: now,
		LastLogin: now,

		ReadingsLists: []models.ReadingsList{},
		Roles:         []string{"common_user"},
		Profile:       models.Profile{},
	}
}

func (um UserMapper) UpdateDtoToEntity(currentUser models.User, userUpdateDTO dtos.UserInsertDTO) models.User {
	currentUser.Email = userUpdateDTO.Email
	currentUser.UpdatedAt = time.Now().UTC().Local()
	currentUser.Username = userUpdateDTO.Username
	currentUser.Password = userUpdateDTO.Password

	return currentUser
}
