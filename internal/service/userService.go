package service

import (
	"fmt"

	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/dto"
	"github.com/hritesh04/url-shortner/internal/helper"
)

type UserRepository interface {
	InsertUser(*dto.SignUpRequest)(int32,error)
	GetUserByEmail(string)(*dto.Users,error)
	GetUserById(int32)(*dto.Users,error)
}

type UserService struct {
	Repo UserRepository
	Auth rest.Auth
}

func (us *UserService)CreateUser(reqForm *dto.SignUpRequest)(string,error){
	if isValid := helper.IsValidEmail(reqForm.Email); !isValid {
		return "",fmt.Errorf("invalid email")
	}
	hashedPassword := us.Auth.HashPassword(reqForm.Password)
	if hashedPassword == "" {
		return "",fmt.Errorf("failed to hash password")
	}
	reqForm.Password = hashedPassword
	userId,err := us.Repo.InsertUser(reqForm)
	if err != nil {
		return "",err
	}
	token,err := us.Auth.GenerateToken(userId,"free")
	if err != nil{
		return "",err
	}
	return token,nil
}

func (us *UserService)Login(reqForm *dto.SignInRequest)(string,error){
	user,err := us.Repo.GetUserByEmail(reqForm.Email)
	if err != nil {
		return "",err
	}
	if validPass := us.Auth.ComparePassword(reqForm.Password,user.Password); !validPass {
		return "",fmt.Errorf("invalid password")
	}
	token,err := us.Auth.GenerateToken(user.Id,user.SubscriptionType)
	if err != nil {
		return "",fmt.Errorf("errror creating token")
	}
	return token,nil
}

func (us *UserService)GetUserDetails(id int32)(*dto.Users,error){
	user,err := us.Repo.GetUserById(id)
	
	if err != nil {
		return user,err
	}

	return user,nil
}