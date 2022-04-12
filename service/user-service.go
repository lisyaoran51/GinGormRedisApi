package service

import (
	

	"github.com/lisyaoran51/GinGormRedisApi/entity"
	"github.com/lisyaoran51/GinGormRedisApi/repository"
)

type UserService interface {
	Save(entity.User) entity.User
	Update(entity.User) error
	Delete(entity.User) error
	FindAll() []entity.User
	FindByID(id int) (entity.User, error)
}

type userService struct{
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) Save(user entity.User) entity.User {
	// service.videos = append(service.videos, video)
	service.userRepository.Save(user)
	return user
}

func (service *userService) Update(user entity.User) error {
	return service.userRepository.Update(user)
}

func (service *userService) Delete(user entity.User) error {
	return service.userRepository.Delete(user)
}

func (service *userService) FindAll() []entity.User {
	return service.userRepository.FindAll()
}

func (service *userService) FindByID(id int) (entity.User, error) {


	return service.userRepository.FindByID(id)
}
