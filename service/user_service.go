package service

import (
	"crypto/sha256"
	"encoding/hex"
	"go_test/repository"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(newUser *UserRequest) (*UserResponse, error) {

	h := sha256.New()
	h.Write([]byte(newUser.Password))
	sha256_hash := hex.EncodeToString(h.Sum(nil))

	userCreate := repository.User_db{
		UserID:           primitive.NewObjectID(),
		Username:         newUser.Username,
		Password:         sha256_hash,
		Firstname:        newUser.Firstname,
		Lastname:         newUser.Lastname,
		Birthdate:        newUser.Birthdate,
		Programing_skill: newUser.Programing_skill,
	}

	user_db, err := s.userRepo.Create(userCreate)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user := UserResponse{
		UserID:           user_db.UserID,
		Username:         user_db.Username,
		Firstname:        user_db.Firstname,
		Lastname:         user_db.Lastname,
		Birthdate:        user_db.Birthdate,
		Programing_skill: user_db.Programing_skill,
	}

	return &user, err
}

func (s *userService) LoginUser(loginUser *UserLogin) (string, error) {

	user := repository.User_db{
		Username: loginUser.Username,
		Password: loginUser.Password,
	}
	getuser, err := s.userRepo.Login(user)
	if err != nil {
		log.Println(err)
		return "", err
	}
	secret_key := []byte("my_secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  getuser.UserID,
		"username": getuser.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString(secret_key)

	return tokenString, err
}

func (s *userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.Getall()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	usersResponse := []UserResponse{}

	for _, user := range users {
		userResponse := UserResponse{
			UserID:           user.UserID,
			Username:         user.Username,
			Firstname:        user.Firstname,
			Lastname:         user.Lastname,
			Birthdate:        user.Birthdate,
			Programing_skill: user.Programing_skill,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse, err
}

func (s *userService) GetUser(user_id string) (*UserResponse, error) {
	user, err := s.userRepo.GetById(user_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := UserResponse{
		UserID:           user.UserID,
		Username:         user.Username,
		Firstname:        user.Firstname,
		Lastname:         user.Lastname,
		Birthdate:        user.Birthdate,
		Programing_skill: user.Programing_skill,
	}

	return &userResponse, err

}

func (s *userService) UpdateUser(user_id string, requestUser *UpdateRequest) (*UserResponse, error) {
	user := repository.User_db{
		Firstname:        requestUser.Firstname,
		Lastname:         requestUser.Lastname,
		Birthdate:        requestUser.Birthdate,
		Programing_skill: requestUser.Programing_skill,
	}
	updatedUser, err := s.userRepo.Update(user_id, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userResponse := UserResponse{
		UserID:           updatedUser.UserID,
		Username:         updatedUser.Username,
		Firstname:        updatedUser.Firstname,
		Lastname:         updatedUser.Lastname,
		Birthdate:        updatedUser.Birthdate,
		Programing_skill: updatedUser.Programing_skill,
	}
	return &userResponse, err
}

func (s *userService) DeleteUser(user_id string) (*DeleteResponse, error) {
	user, err := s.userRepo.Delete(user_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	deleteUser := DeleteResponse{
		UserID:   user.UserID,
		Username: user.Username,
	}
	return &deleteUser, err
}
