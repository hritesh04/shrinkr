package routes

import (
	"expvar"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/models"
	"github.com/hritesh04/url-shortner/prometheus"
)


var stats = expvar.NewMap("urlVisitCount").Init()

func Resolve(c *fiber.Ctx)error {
	url := c.Params("url")

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
	prometheus.UrlVisitCount.WithLabelValues(url).Inc()
	return c.Redirect(urlDetails.Original,302)
}

func GetStats(c *fiber.Ctx)error{
	url := c.Query("url")
	value := stats.Get(url).(*expvar.Int).Value()
	return c.Status(200).JSON(&fiber.Map{
		"success":true,
		"data":value,
	})
}