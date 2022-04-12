package repository

import (
	"github.com/lisyaoran51/GinGormRedisApi/entity"
)


type UserRepository interface {
	Save(user entity.User)
	Update(user entity.User) error
	Delete(user entity.User) error
	FindAll() []entity.User
	FindByID(id int) (entity.User, error)
}
