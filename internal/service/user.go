package service

import (
	"clean-API/internal/dto"
	"clean-API/internal/model"
	"clean-API/internal/repository"
	"clean-API/internal/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//go:generate mockgen -source=user.go -destination=user_mock.go -package service

type UserService interface {
	Login(email, password string) (string, error)
	Register(email, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	config         dto.Config
}

func newUserService(userRepository repository.UserRepository, config dto.Config) UserService {
	return &userService{userRepository, config}
}

func (u userService) Login(email, password string) (string, error) {
	user, err := u.userRepository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	if ok := utils.CheckPasswordHash(password, user.Password); !ok {
		return "", fmt.Errorf("invalid password")
	}

	return u.generateJwt(user)
}

func (u userService) Register(email, password string) (string, error) {

	if email == "" {
		return "", fmt.Errorf("email is required")
	}

	if password == "" {
		return "", fmt.Errorf("password is required")
	}

	if len(password) < 3 {
		return "", fmt.Errorf("password must be at least 8 characters")
	}

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return "", fmt.Errorf("error hashing password")
	}

	userByEmail, err := u.userRepository.GetUserByEmail(email)

	if userByEmail.ID != 0 {
		return "", fmt.Errorf("user with email %s already exists", email)
	}
	user := model.User{Email: email, Password: hashedPassword}

	user, err = u.userRepository.Save(user)

	if err != nil {
		return "", fmt.Errorf("error saving user")
	}

	return u.generateJwt(user)
}

func (u userService) generateJwt(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(u.config.SigningSecret))
}
