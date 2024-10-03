package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
)

type AuthMapper struct {
}

func (am AuthMapper) SignupDtoEntity(signupDTO dtos.SignupDTO) models.User {
	firstName := signupDTO.ProprofileInsertDTO.FirstName
	lastname := signupDTO.ProprofileInsertDTO.LastName

	now := time.Now().UTC()
	hashedPassword, _ := utils.HashPassword(signupDTO.Password)
	profile := models.Profile{
		FullName:        firstName + " " + lastname,
		Biography:       signupDTO.ProprofileInsertDTO.Biography,
		ProfileCoverURL: signupDTO.ProprofileInsertDTO.ProfileCoverURL,
		ProfileImageURL: signupDTO.ProprofileInsertDTO.ProfileImageURL,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	return models.User{
		Username:  signupDTO.Username,
		Email:     signupDTO.Email,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
		LastLogin: now,

		Roles:   []string{"common_user"},
		Profile: profile,
	}

}
