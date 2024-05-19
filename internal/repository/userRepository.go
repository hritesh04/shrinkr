package repository

import (
	"context"
	"database/sql"

	"github.com/hritesh04/url-shortner/internal/dto"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB)*UserRepository{
	return&UserRepository{
		DB:	db,
	}
}

func (ur * UserRepository)InsertUser(formData *dto.SignUpRequest)(int32,error){
	var userId int32
	ctx := context.Background()

	err := ur.DB.QueryRowContext(ctx,"INSERT INTO USERS (name,email,password) VALUES ($1,$2,$3) returning id;",formData.Name,formData.Email,formData.Password).Scan(&userId)

	if err != nil {
		return userId,err
	}
	
	return userId,nil
}

func (ur *UserRepository)GetUserByEmail(email string)(*dto.Users,error){
	ctx := context.Background()
	var user dto.Users
	rows,err := ur.DB.QueryContext(ctx,"SELECT * FROM USERS WHERE email=$1",email)
	if err != nil {
		return &dto.Users{},err
	}

	defer rows.Close()
	
	for rows.Next(){
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.SubscriptionType)
		if err != nil{
			return &dto.Users{},err
		}
	}

	return &user,nil
}

func (ur *UserRepository)GetUserById(id int32)(*dto.Users,error){
	ctx:=context.Background()
	
	var user dto.Users

	rows,err := ur.DB.QueryContext(ctx,"SELECT u.id,u.name,u.email,u.password,u.subscription_type,url.id,url.original,url.shortened,url.user_id,url.rateremaining,url.expiry,url.isactive FROM USERS u LEFT JOIN URLS url ON u.id = url.user_id WHERE u.id=$1",id)
	
	if err !=nil {
		return nil,err
	}

	defer rows.Close()

	for rows.Next(){
		var url dto.Url
		
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.SubscriptionType,&url.Id,&url.Original,&url.Shortened,&url.User_id,&url.RateRemaining,&url.Expiry,&url.IsActive)

		if err !=nil{
			return &user,nil
		}

		user.Urls = append(user.Urls, url)
	}

	if err := rows.Err(); err != nil {
        return &user, err
    }

    return &user, nil
}