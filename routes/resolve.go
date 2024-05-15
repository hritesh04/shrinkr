package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/models"
	"github.com/hritesh04/url-shortner/prometheus"
)

func Resolve(c *fiber.Ctx)error {
	url := c.Params("url")

	ctx := context.Background()

	val,err := database.Cache.Get(ctx,url).Result()
	if err == nil {
		prometheus.UrlVisitCount.WithLabelValues(url).Inc()
        return c.Redirect(val,302)
    }
	urlDetails:= models.Url{}
	
	db := database.Connect()
	defer db.Close()

	rows,err:=db.Query("SELECT original FROM urls WHERE shortened = $1",url)
	if err != nil{
		fmt.Println(err)
		return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
			"success":false,
			"data":"Invalid url",
		})
	}
	for rows.Next(){
		err := rows.Scan(&urlDetails.Original)
		if err != nil{
			fmt.Println(err)
			return c.Status(400).JSON(&fiber.Map{
				"success":false,
				"data":"Error scaning data",
			})
		}
	}
	cacheErr := database.Cache.Set(ctx,url,urlDetails.Original,time.Minute*10).Err()
	if cacheErr != nil {
		fmt.Println(cacheErr)
	}
	prometheus.UrlVisitCount.WithLabelValues(url).Inc()
	return c.Redirect(urlDetails.Original,302)
}