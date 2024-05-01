package helper

import (
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/models"
)

func AddUrl(urlData *models.Url)(string,error){
	db := database.Connect()
	defer db.Close()
	var shortenedUrl string
	err := db.QueryRow("INSERT INTO URLS (original,shortened,user_id) VALUES ($1,$2,$3) RETURNING shortened;", urlData.Original,urlData.Shortened,urlData.User_id).Scan(&shortenedUrl)
	if err != nil {
		return "",err
	}
	return shortenedUrl,nil
}