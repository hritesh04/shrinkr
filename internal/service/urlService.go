package service

import (
	"fmt"
	"os"
	"time"

	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/dto"
)

type UrlRepository interface {
	AddUrl(*dto.Request,int32,int32,time.Time)(*dto.Url,error)
	DeleteUrl()
	UpdateUrl()
	FindUrlMetricsById()
	FindUrlByUser()
	Resolve(string)(string,error)
	GetCache(string)(string,error)
	SetCache(string,string,time.Duration)error
}

type UrlService struct {
	Repo 	UrlRepository
	Auth 	rest.Auth
	Monitor rest.Monitor
}

func (u *UrlService)Resolve(url string)(string,error){
	val,_ := u.Repo.GetCache(url)
	if val != "" {
		u.Monitor.GetCounter("UrlVisitCount").WithLabelValues(url).Inc()
		return val,nil
	}
	original,err:= u.Repo.Resolve(url)
	if err != nil {
		return "",err
	}
	fmt.Printf("here : %v",original)
	cacheErr := u.Repo.SetCache(url,original,time.Minute*10)
	if cacheErr != nil {
		return "",cacheErr
	}
	u.Monitor.GetCounter("UrlVisitCount").WithLabelValues(url).Inc()

	return original,nil
}

func (u *UrlService)ShortenUrl(url *dto.Request,user *dto.Token)(*dto.Url,error){
	var newUrl *dto.Url
	var err error = nil
	if user.SubscriptionType == os.Getenv("SUB_PRE") {
		var rate int32 = 100000
		expiry := time.Now().Add(30*24*time.Hour)
		newUrl,err = u.Repo.AddUrl(url,user.Id,rate,expiry)
		if err != nil {
			return newUrl,err
		}
		return newUrl,err
	}
	expiry := time.Now().Add(7*24*time.Hour)
	newUrl,err = u.Repo.AddUrl(url,user.Id,1000,expiry)
	if err != nil {
		return newUrl,err
	}
	return newUrl,err
}