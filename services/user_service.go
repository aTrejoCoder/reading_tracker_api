package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUserById(userId primitive.ObjectID) (*dtos.UserDTO, error)
	CreateUser(userInsertDTO dtos.UserInsertDTO) (*dtos.UserDTO, error)
	UpdateUser(userId primitive.ObjectID, userUpdateDTO dtos.UserInsertDTO) (*dtos.UserDTO, error)
	DeleteUser(userId primitive.ObjectID) (*dtos.UserDTO, error)
}

type userServiceImpl struct {
	userMapper       mappers.UserMapper
	userRepository   repository.UserExtendRepository
	commonRepository repository.Repository[models.User]
}

func NewUserService(commonRepository repository.Repository[models.User],
	userRepository repository.UserExtendRepository) UserService {
	return &userServiceImpl{
		commonRepository: commonRepository,
		userRepository:   userRepository,
	}
}

func (us userServiceImpl) GetUserById(userId primitive.ObjectID) (*dtos.UserDTO, error) {
	user, err := us.commonRepository.GetByID(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	userDTO := us.userMapper.EntityToDTO(*user)
	return &userDTO, nil
}

func (us userServiceImpl) CreateUser(userInsertDTO dtos.UserInsertDTO) (*dtos.UserDTO, error) {
	newUser := us.userMapper.InsertDtoToEntity(userInsertDTO)

	_, err := us.commonRepository.Create(context.TODO(), &newUser)
	if err != nil {
		return nil, err
	}

	userFounded, err := us.userRepository.GetByEmail(context.TODO(), newUser.Email)
	if err != nil {
		return nil, err
	}

	userDTO := us.userMapper.EntityToDTO(*userFounded)
	return &userDTO, nil
}

func (us userServiceImpl) UpdateUser(userId primitive.ObjectID, userUpdateDTO dtos.UserInsertDTO) (*dtos.UserDTO, error) {
	currentUser, err := us.commonRepository.GetByID(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	userUpdated := us.userMapper.UpdateDtoToEntity(*currentUser, userUpdateDTO)

	if _, err := us.commonRepository.Update(context.TODO(), userUpdated.Id, userUpdated); err != nil {
		return nil, err
	}

	userDTO := us.userMapper.EntityToDTO(userUpdated)
	return &userDTO, nil
}

func (us userServiceImpl) DeleteUser(userId primitive.ObjectID) (*dtos.UserDTO, error) {
	user, err := us.commonRepository.GetByID(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	us.commonRepository.DeleteById(context.TODO(), userId)

	userDTO := us.userMapper.EntityToDTO(*user)
	return &userDTO, nil
}
