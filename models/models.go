package models

import (
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"
)

type Users struct {
	Id       int32
	Name     string
	Email    string
	Password string
	Urls     []Url
}

type Url struct {
	Id             int32
	Original       string
	Shortened      string
	User_id        int32
	RateRemaining  int32
	Expiry         time.Time
	RateLimitReset time.Time
	IsActive       bool
}

type Claim struct {
	jtoken.RegisteredClaims
	Id int32
}

type SignUpRequest struct {
	Name string
	Email string
	Password string
}

type Request struct{
	Url string	
	CustomUrl string
}