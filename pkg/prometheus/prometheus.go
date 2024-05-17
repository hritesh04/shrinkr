package prometheus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/internal/dto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type MonitorService struct{
	Topics map[string]*prometheus.CounterVec
}

// var UrlVisitCount = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "url_visit_count",
// 		Help: "Total number of time URL is visited",
// 	},
// 	[]string{"url"},
// )

// .\prometheus.exe --storage.tsdb.retention.time=30d  for rentiontime

func (m *MonitorService)Metrics(c *fiber.Ctx)error {
	adapter := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	adapter(c.Context())
	return nil
}

func (m *MonitorService)GetCounter(name string)*prometheus.CounterVec{
    return m.Topics[name]
}

func InitMetrics(topics map[string]*prometheus.CounterVec)*MonitorService {
	fmt.Println("Registered")
	for _,topic := range topics {
		prometheus.MustRegister(topic)
	}
	return &MonitorService{
		Topics: topics,
	}
}


func (m *MonitorService) GetStats(topic, step, limit string) (dto.QueryResponse, error) {
    // Set the time range based on the limit
    offset, err := time.ParseDuration(limit)
    if err != nil {
        return dto.QueryResponse{}, err
    }
    now := time.Now().UTC()
    end := now
    start := now.Add(-offset)

    // Construct the Prometheus query
    query := fmt.Sprintf("sum(url_visit_count{job=\"prometheus\",url=\"%s\"})&start=%s&end=%s&step=%s",
        topic, start.Format(time.RFC3339), end.Format(time.RFC3339), step)

    // Construct the URL for the Prometheus query endpoint
    url := fmt.Sprintf("http://localhost:9090/api/v1/query?query=%s", query)

    // Send the HTTP request to the Prometheus query endpoint
    client := http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return dto.QueryResponse{}, err
    }
    res, err := client.Do(req)
    if err != nil {
        return dto.QueryResponse{}, err
    }
    defer res.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return dto.QueryResponse{}, err
    }

    // Parse the response body as JSON
    var response dto.QueryResponse
    if err := json.Unmarshal(body, &response); err != nil {
        return dto.QueryResponse{}, err
    }

    return response, nil
}