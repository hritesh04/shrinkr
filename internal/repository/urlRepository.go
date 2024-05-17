package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/dto"
)

type UrlRepository struct {
	DB *sql.DB
	Cache rest.Cache
}

func NewUrlRepository(db *sql.DB,Cache rest.Cache) *UrlRepository {
	return &UrlRepository{
		DB: db,
		Cache: Cache,
	}
}

func (r *UrlRepository) AddUrl(url *dto.Request,userId int32,rate int32, expiry time.Time)(*dto.Url,error){
	ctx := context.Background()
	urlData := dto.Url{}

	err := r.DB.QueryRowContext(ctx,"INSERT INTO URLS (original,shortened,user_id,rateremaining,expiry) VALUES ($1,$2,$3,$4,$5) RETURNING *;", url.Url,url.CustomUrl,userId,rate,expiry).Scan(&urlData)
	if err != nil {
		return &urlData,err
	}
	return &urlData,nil

}
func (r *UrlRepository) DeleteUrl(){

}
func (r *UrlRepository) UpdateUrl(){

}
func (r *UrlRepository) FindUrlByUser(){

}
func (r *UrlRepository) FindUrlMetricsById(){

}
func (r *UrlRepository) Resolve(short string)(string,error){
	var original string
	rows,err:=r.DB.Query("SELECT original FROM urls WHERE shortened = $1",short)
	if err != nil{
		fmt.Println(err)
		return "",err
	}
	for rows.Next(){
		err := rows.Scan(&original)
		if err != nil{
			return "",err
		}
	}
	fmt.Println(original)
	return original,nil
}

func (r *UrlRepository) GetCache(url string)(string,error){
	val,err := r.Cache.Get(url)
	if err != nil{
		return "",err
	}
	return val,nil
}
func (r *UrlRepository) SetCache(key string ,value string, time time.Duration)error{
	ctx := context.Background()
	err := r.Cache.Set(ctx,key,value,time); if err != nil{
		return err
	}
	return nil
}