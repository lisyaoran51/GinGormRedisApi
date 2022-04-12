package repository

import (
	// "database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/lisyaoran51/GinGormRedisApi/entity"


)



type mysqlUserRepository struct {
	connection *gorm.DB
}


func NewMysqlUserRepository(dsn string) UserRepository {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	return &mysqlUserRepository {
		connection: db,
	}
}

func (db *mysqlUserRepository) Save(user entity.User) {
	user.CreatedAt = time.Now()
	fmt.Println("mysqlUserRepository.Save():", user)

	db.connection.Create(&user)

	// db.connection.Update("createdAt", time.Now())
}

func (db *mysqlUserRepository) Update(user entity.User) error {

	fmt.Println("mysqlUserRepository.Update():", user)

	// var users []entity.User

	// db.connection.Find(&users)
	// fmt.Println("all user count:", len(users))
	// for _, v := range users {
    //     fmt.Println("user:", v)
    // }

	// fmt.Println("===============")

	// db.connection.Where("userId = ?", user.UserId).Find(&users)
	
	// for _, v := range users {
    //     fmt.Println("user:", v)
    // }

	err := db.connection.First(&user, user.UserId).Error
	if err != nil {
		fmt.Println("mysqlUserRepository.Update(): no such user with id ", user.UserId)
		return err
	}

	err = db.connection.Updates(&user).Error
	// err := db.connection.Where("userId = ?", user.UserId).Updates(&user).Error
	// err := db.connection.Where("account = ?", user.Account).Updates(&user).Error
	// err := db.connection.Updates(&user).Error
	if err != nil{
		fmt.Println("error:", err)
	}

	return err
}

func (db *mysqlUserRepository) Delete(user entity.User) error {

	fmt.Println("mysqlUserRepository.Delete():", user)

	err := db.connection.First(&user, user.UserId).Error
	if err != nil {
		fmt.Println("mysqlUserRepository.Delete(): no such user with id ", user.UserId)
		return err
	}

	db.connection.Delete(&user)

	return nil
}

func (db *mysqlUserRepository) FindAll() []entity.User {
	var users []entity.User

	db.connection.Set("gorm:auto_preload", true).Find(&users)
	return users
}

func (db *mysqlUserRepository) FindByID(id int) (entity.User, error) {
	var user entity.User
	err := db.connection.First(&user, id).Error


	if err != nil {
		return user, err
	}

	return user, nil
}