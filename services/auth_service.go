package services

import (
	"context"
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/middleware/token"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
)

type AuthServices interface {
	ValidateSignupCredentials(signupDTO dtos.SignupDTO) error
	ProccesSignup(signupDTO dtos.SignupDTO) (string, error)
	ValidateLoginCredentials(loginDTO dtos.LoginDTO) (*dtos.UserDTO, error)
	ProccesLogin(userDTO dtos.UserDTO) (string, error)
}

func NewAuthService(userRepository repository.UserExtendRepository, commonRepository repository.Repository[models.User]) AuthServices {
	return &authServiceImpl{
		userRepository:   userRepository,
		commonRepository: commonRepository,
	}
}

type authServiceImpl struct {
	userRepository   repository.UserExtendRepository
	authMapper       mappers.AuthMapper
	userMapper       mappers.UserMapper
	commonRepository repository.Repository[models.User]
}

func (as authServiceImpl) ValidateSignupCredentials(signupDTO dtos.SignupDTO) error {
	_, err := as.userRepository.GetByEmail(context.Background(), signupDTO.Email)
	if err != nil && err != utils.ErrNotFound {
		return utils.ErrDatabase
	}

	if err == nil {
		return errors.New("email is already in use")
	}

	_, err = as.userRepository.GetByUsername(context.Background(), signupDTO.Username)
	if err != nil && err != utils.ErrNotFound {
		return utils.ErrDatabase
	}

	if err == nil {
		return errors.New("username is already in use")
	}

	return nil
}

func (as authServiceImpl) ProccesSignup(signupDTO dtos.SignupDTO) (string, error) {
	newUser := as.authMapper.SignupDtoEntity(signupDTO)

	if _, err := as.commonRepository.Create(context.Background(), &newUser); err != nil {
		return "", err
	}

	userCreated, err := as.userRepository.GetByUsername(context.Background(), signupDTO.Username)
	if err != nil {
		return "", err
	}

	jwt, err := token.GenerateJWT(userCreated.Id.Hex(), newUser.Username, newUser.Email, []string{"common_user"})
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (as authServiceImpl) ValidateLoginCredentials(loginDTO dtos.LoginDTO) (*dtos.UserDTO, error) {
	var userFounded *models.User
	var err error

	userFounded, err = as.userRepository.GetByEmail(context.Background(), loginDTO.LoginIdentfier)
	if err == utils.ErrNotFound {
		userFounded, err = as.userRepository.GetByUsername(context.Background(), loginDTO.LoginIdentfier)
	}

	if err != nil || userFounded == nil {
		return nil, errors.New("user not found with given credentials")
	}

	isPasswordCorrect := utils.ComparePasswords(userFounded.Password, loginDTO.Password)
	if !isPasswordCorrect {
		return nil, utils.ErrUnauthorized
	}

	userDTO := as.userMapper.EntityToDTO(*userFounded)
	return &userDTO, nil
}
func (as authServiceImpl) ProccesLogin(userDTO dtos.UserDTO) (string, error) {
	errChan := make(chan error)
	jwtChan := make(chan string)

	// Update Last Login
	go func() {
		if err := as.userRepository.UpdateLastLogin(context.Background(), userDTO.Id); err != nil {
			errChan <- err
		}
		errChan <- nil
	}()

	// Generate JWT
	go func() {
		jwt, err := token.GenerateJWT(userDTO.Id.Hex(), userDTO.Username, userDTO.Email, userDTO.Roles)
		if err != nil {
			errChan <- err
		}
		jwtChan <- jwt
	}()

	for i := 0; i < 2; i++ {
		select {
		case err := <-errChan:
			if err != nil {
				return "", err
			}
		case jwt := <-jwtChan:
			return jwt, nil
		}
	}

	return "", nil
}
