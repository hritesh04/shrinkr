package prometheus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var UrlVisitCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "url_visit_count",
		Help: "Total number of time URL is visited",
	},
	[]string{"url"},
)

// .\prometheus.exe --storage.tsdb.retention.time=30d  for rentiontime

func Metrics(c *fiber.Ctx)error {
	adapter := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	adapter(c.Context())
	return nil
}

func InitMetrics() {
	fmt.Println("Registered")
	prometheus.MustRegister(UrlVisitCount)
}


func GetStats(c *fiber.Ctx)error{
	// Set the time range to 7 days
    now := time.Now().UTC()
    end := now //time.Now()

	// get from the time the url was created
    start := now.Add(-time.Second*600)
	step:= c.Query("step") // according to the query

	// url get from the query as well
    query := fmt.Sprintf("sum(url_visit_count{job=\"prometheus\",url=\"draw\"})&start=%s&end=%s&step=%s", start.Format(time.RFC3339), end.Format(time.RFC3339),step)
    
	
	url := fmt.Sprintf("http://localhost:9090/api/v1/query?query=%s",query)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"data":err,
		})
	}

	client := &http.Client{}

	res,err := client.Do(req)
    if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success":false,
			"data":err,
		})
    }

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"data":err,
		})
	}
        // Parse the response body to JSON
        var data map[string]interface{} 
        if err := json.Unmarshal(body, &data); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
                "success": false,
                "error":   fmt.Sprintf("Error parsing response body to JSON: %v", err),
            })
        }

	return c.Status(200).JSON(data)
}
