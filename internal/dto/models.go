package dto

import (
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"
)

type Users struct {
	Id       			int32	`json:"id"`
	Name     			string	`json:"name"`
	Email    			string	`json:"email"`
	Password 			string	`json:"password"`
	SubscriptionType	string 	`json:"subscriptionType"`
	Urls     			[]Url	`json:"urls"`
}

type Url struct {
	Id             int32		`json:"id"`
	Original       string		`json:"original"`
	Shortened      string		`json:"shortened"`
	User_id        int32		`json:"userId"`
	RateRemaining  int32		`json:"rateRemaining"`
	Expiry         time.Time	`json:"expiry"`
	IsActive       bool			`json:"isActive"`
}

type Claim struct {
	jtoken.RegisteredClaims
	Id int32
	SubscriptionType string
}

type Token struct {
	Id int32
	SubscriptionType string
}

type SignUpRequest struct {
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}
type SignInRequest struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type Request struct{
	Url			string	`json:"url"`
	CustomUrl	string	`json:"customUrl"`
}

type QueryResponse struct {
        ResultType string `json:"resultType"`
        Result     []struct {
            Metric map[string]string `json:"metric"`
            Values []interface{}     `json:"values"`
    } `json:"result"`
}