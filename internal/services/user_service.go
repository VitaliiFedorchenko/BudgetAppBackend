package services

import (
	"BudgetApp/cmd/server/validation"
	"BudgetApp/internal/configs"
	"BudgetApp/internal/utils"
	"BudgetApp/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	return &UserService{db: db}
}

func (s *UserService) CreateUser(req validation.CreateUserRequest) (*models.User, error) {
	var user models.User
	var err error

	user.Email = req.Email
	user.Name = req.Name
	user.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetUserViaEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("there is no user with such email")
	}

	return &user, nil
}
